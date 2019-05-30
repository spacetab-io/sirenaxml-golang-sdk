package client

import (
	"bufio"
	"net"
	"time"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
)

// Packet represents application level data.
type Packet struct {
	header           *Header
	subHeader        []byte
	message          []byte
	messageSignature []byte
}

var (
	respPool *RespPool
	msgPool  *MsgPool
)

// Channel wraps user connection.
type Channel struct {
	cfg    *sirenaXML.Config
	conn   net.Conn // Socket connection.
	send   chan *Packet
	Key    []byte
	KeyID  uint32
	Logger logs.LogWriter
}

func NewChannel(sc *sirenaXML.Config) (*Channel, error) {
	err := sc.PrepareKeys()
	if err != nil {
		return nil, err
	}
	conn, err := net.Dial("tcp", sc.GetAddr())
	if err != nil {
		return nil, errors.Wrap(err, "dial sirena addr error")
	}

	c := &Channel{
		conn: conn,
		send: make(chan *Packet, sc.RequestHandlers),
		cfg:  sc,
	}

	respPool = NewRespPool()
	msgPool, err = NewMsgPool(respPool, sc.RequestHandlers)
	if err != nil {
		return nil, err
	}

	err = createSignKey(c)
	if err != nil {
		return nil, err
	}

	go func() {
		buf := bufio.NewReader(c.conn)
		for {
			err := c.readPacket(buf)
			if err != nil {
				c.Logger.Error(err) // log it for now @TODO change it
			}
		}
	}()

	return c, err
}

func (c *Channel) SendMsg(msg []byte) ([]byte, error) {
	p, err := c.NewRequest(msg)
	if err != nil {
		return nil, err
	}
	c.sendPacket(p)

	response := getResponseFromMsgPool(p.header.MessageID)
	return response.message, nil
}

func (c *Channel) sendPacket(p *Packet) {
	buf := bufio.NewWriter(c.conn)

	if err := writePacket(buf, p); err != nil {
		panic(err) // panic for now @TODO change it
	}
	_ = buf.Flush()
}

func (c *Channel) SetLogger(l logs.LogWriter) {
	c.Logger = l
}

func createSignKey(c *Channel) error {
	// Create symmetric key
	if err := c.signKey(); err != nil {
		return errors.Wrap(err, "creating and signing key error")
	}
	// Update key every 1 hour
	go func() {
		for range time.Tick(time.Hour) {
			if err := c.signKey(); err != nil {
				logs.Logger.Fatal("key updating error")
			}
		}
	}()

	return nil
}

func getResponseFromMsgPool(msgID uint32) *Packet {
	resp := respPool.GetPacket(msgID)
	msgPool.ReturnMsgIDToPool(msgID)
	return resp
}
