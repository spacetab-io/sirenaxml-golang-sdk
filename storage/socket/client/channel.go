package client

import (
	"bufio"
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"io"
	"net"
	"sync"
	"time"
)

// Packet represents application level data.
type Packet struct {
	header           *Header
	subHeader        []byte
	message          []byte
	messageSignature []byte
}

type Config struct {
	Ip                       string
	Environment              string
	ClientPublicKey          string
	ClientPrivateKey         string
	ClientPrivateKeyPassword string
	ServerPublicKey          string
	Address                  string
	ClientID                 uint16
	MaxConnections           uint32
	ZippedMessaging          bool
	MaxConnectTries          int
	Buffer                   int
}

var (
	respPool *RespPool
	msgPool  *MsgPool
)

//Channel wraps user connection.
type Channel struct {
	cfg    Config
	send   chan *Packet
	socket *socket
	Logger logs.LogWriter
}

type KeyData struct {
	ID  uint32 `json:"id"`
	Key []byte `json:"key"`
}

type socket struct {
	m        sync.Mutex
	KeyData  KeyData
	cancel   context.CancelFunc
	conn     net.Conn  // Socket connection session
	addr     string    // address
	sessNum  int       // number of success tries
	initTime time.Time // time of last connection
}

const (
	EnvLearning   = "GRU"
	EnvTesting    = "GRT"
	EnvProduction = "GRS"
)

var (
	portsMap = map[string]string{
		EnvLearning:   "34323",
		EnvTesting:    "34322",
		EnvProduction: "34321",
	}
)

// GetAddr return sirena address to connect client to
func (config *Config) GetAddr() (string, error) {
	return config.Ip + ":" + portsMap[config.Environment], nil
}

func NewChannel(l logs.LogWriter, opts ...Option) (*Channel, error) {

	respPool = NewRespPool()

	c := &Channel{
		send: make(chan *Packet, 0),
	}

	c.SetLogger(l)

	for _, opt := range opts {
		opt(c)
	}

	addr, err := c.cfg.GetAddr()

	spew.Dump(addr, "========-----------============")

	if err != nil {
		return nil, err
	}

	c.socket = &socket{addr: addr}
	msgPool, err = NewMsgPool(respPool, c.cfg.MaxConnections)
	err = tryToConnect(c)

	return c, err
}

func tryToConnect(c *Channel) (err error) {

	//spew.Dump("===================", c.cfg.MaxConnectTries)
	for i := 0; i <= c.cfg.MaxConnectTries; i++ {
		c.Logger.Debugf("connection try %d start", i)

		err = c.connect()
		spew.Dump("connection error: ", err)
		if err == nil {

			c.Logger.Debugf("connection try %d succeed", i)
			break
		}
	}
	return err
}

func (c *Channel) connect() error {
	conn, err := net.DialTimeout("tcp", c.socket.addr, 1*time.Second)
	if err != nil {
		return errors.Wrap(err, "dial sirena addr error")
	}

	c.socket.conn = conn
	err = c.signKey()
	if err != nil {
		return err
	}

	c.socket.sessNum++
	c.socket.initTime = time.Now()
	c.Logger.Infof("started connection session number %d", c.socket.sessNum)

	var ctx context.Context

	ctx, c.socket.cancel = context.WithCancel(context.Background())

	// start listener
	go func(ctx context.Context) {
		c.Logger.Infof("listening session %d", c.socket.sessNum)
		buf := bufio.NewReader(c.socket.conn)
		var err error

	incomingDataReader:
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				return
			default:
				err = c.readPacket(buf)
				if err != nil && err == io.EOF {
					break incomingDataReader
				} else if err != nil {
					c.Logger.Errorf("reading packet error: %v", err)
				}
			}
		}
		c.reconnect(err)
	}(ctx)

	// Update symmetric key every 1 hour
	go func(ctx context.Context) {
		for range time.Tick(time.Hour) {
			ctx.Done()
		}
	}(ctx)
	return err
}

func (c *Channel) stopListener() {
	c.socket.cancel() // close listener goroutine
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
	c.socket.m.Lock()
	c.clearConnect()
	now := time.Now()
	trottlingLimit := c.socket.initTime.Add(1 * time.Second)
	if now.Sub(trottlingLimit) < 0 {
		panic(errors.New("stop dosing sirena socket"))
	}
	c.Logger.Warningf("reconnect! %s", err)
	err = tryToConnect(c)
	if err != nil {
		panic(err)
	}
	c.socket.m.Unlock()
}

func (c *Channel) disconnect() error {
	return c.socket.conn.Close()
}

func (c *Channel) GetKeyData() KeyData {
	return c.socket.KeyData
}

func (c *Channel) clearConnect() {
	if c.socket.conn != nil {
		_ = c.socket.conn.Close()
	}
	c.socket.conn = nil
	c.socket.KeyData.Key = nil
	c.socket.KeyData.ID = 0
}

func getResponseFromMsgPool(msgID uint32) *Packet {
	resp := respPool.GetPacket(msgID)
	msgPool.ReturnMsgIDToPool(msgID)
	return resp
}
