package client

import (
	"bufio"
	"context"
	"io"
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
	send   chan *Packet
	socket *socket
	Logger logs.LogWriter
}

type KeyData struct {
	ID  uint32
	Key []byte
}

type socket struct {
	KeyData  KeyData
	cancel   context.CancelFunc
	conn     net.Conn  // Socket connection session
	addr     string    // address
	sessNum  int       // number of success tries
	initTime time.Time // time of last connection
}

func NewChannel(sc *sirenaXML.Config, l logs.LogWriter) (*Channel, error) {
	if err := sc.PrepareKeys(); err != nil {
		return nil, err
	}
	addr, err := sc.GetAddr()
	if err != nil {
		return nil, err
	}

	respPool = NewRespPool()
	msgPool, err = NewMsgPool(respPool, sc.MaxConnections)
	if err != nil {
		return nil, err
	}

	c := &Channel{
		socket: &socket{addr: addr},
		send:   make(chan *Packet, sc.MaxConnections),
		cfg:    sc,
	}
	c.SetLogger(l)

	err = c.connect()

	return c, err
}

func (c *Channel) connect() error {
	conn, err := net.Dial("tcp", c.socket.addr)
	if err != nil {
		return errors.Wrap(err, "dial sirena addr error")
	}
	c.socket.conn = conn
	err = createSignKey(c)
	if err != nil {
		return err
	}

	c.socket.sessNum++
	c.socket.initTime = time.Now()
	c.Logger.Infof("started connection session number %d", c.socket.sessNum)
	var ctx context.Context
	ctx, c.socket.cancel = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		c.Logger.Infof("listening session %d", c.socket.sessNum)
		buf := bufio.NewReader(c.socket.conn)
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				break
			default:
				err := c.readPacket(buf)
				if err != nil && err == io.EOF {
					go c.reconnect(err)
					break
				} else if err != nil {
					c.Logger.Errorf("reading packet error: %v", err)
				}
			}
		}
	}(ctx)

	return err
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
	buf := bufio.NewWriter(c.socket.conn)

	if err := writePacket(buf, p); err != nil {
		panic(err) // panic for now @TODO change it
	}
	_ = buf.Flush()
}

func (c *Channel) SetLogger(l logs.LogWriter) {
	c.Logger = l
}

func (c *Channel) reconnect(err error) {
	c.socket.cancel() // close listener goroutine
	c.socket.KeyData.Key = nil
	c.socket.KeyData.ID = 0
	now := time.Now()
	trottlingLimit := c.socket.initTime.Add(1 * time.Second)
	if now.Sub(trottlingLimit) < 0 {
		panic(errors.New("stop dosing sirena socket"))
	}
	c.Logger.Warningf("reconnect! %s", err)
	err = c.connect()
	if err != nil {
		panic(err)
	}
}

func (c *Channel) disconnect() error {
	return c.socket.conn.Close()
}

func (c *Channel) GetKeyData() KeyData {
	return c.socket.KeyData
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
