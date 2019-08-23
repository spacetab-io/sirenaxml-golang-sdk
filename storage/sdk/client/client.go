package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/message"
)

// CustomDialer used to specify dialer.
type CustomDialer interface {
	DialContext(ctx context.Context, network, addr string) (net.Conn, error)
}

// Client wraps user connection.

type Client struct {
	Logger logs.LogWriter

	crypt *sirenaXML.KeysData

	conn *connection

	msgPool  chan uint32
	incoming map[uint32]chan *message.ReceivedMessage
	outgoing chan sending
}

type connection struct {
	clientID          uint16
	KeyData           *message.KeyData
	addr              *net.TCPAddr // address
	dialer            CustomDialer
	sessNum           uint      // number of success tries
	initTime          time.Time // time of last connection
	reconnectAttempts uint
	delayFn           DelayFunc
	useZip            bool
	maxConnections    uint32
}

var (
	// ErrMaxAttemptsExceeded indicates that max reconnect attempts exceeded
	// during connection establishing.
	ErrMaxAttemptsExceeded = errors.New("maximum attempts exceeded")
	errConnClosed          = errors.New("connection closed")
)

// New returns new instance of Client.
func New(opts ...Option) (*Client, error) {
	var err error
	cl := &Client{
		Logger:   logs.NewNullLog(),
		conn:     &connection{maxConnections: 100, KeyData: &message.KeyData{}},
		incoming: make(map[uint32]chan *message.ReceivedMessage),
	}

	for _, opt := range opts {
		opt(cl)
	}

	cl.msgPool = make(chan uint32, cl.conn.maxConnections)
	cl.outgoing = make(chan sending, cl.conn.maxConnections)

	err = cl.generateMsgPool()
	if err != nil {
		return nil, err
	}

	return cl, err
}

func (cl *Client) generateMsgPool() error {
	if cl.conn.maxConnections == 0 {
		return errors.New("empty message pool is not allowed")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cap(cl.msgPool); i++ {
		msgID := r.Uint32()
		cl.incoming[msgID] = make(chan *message.ReceivedMessage, 1)
		cl.msgPool <- msgID
	}

	return nil
}

// Connect will try to keep connection active
// by reconnection on any errors.
// It will return error only if ctx was done.
//
// NOTE: It's blocking method.
// nolint: interfacer
func (cl *Client) Connect(ctx context.Context) error {
	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return cl.dial(gctx)
	})
	g.Go(func() error {
		return cl.signKey(gctx)
	})
	return g.Wait()
}

// establishConn used to establish connection and reconnect on errors.
func (cl *Client) establishConn(ctx context.Context) (net.Conn, error) {
	for attempt := uint(0); attempt <= cl.conn.reconnectAttempts; attempt++ {
		conn, err := cl.conn.dialer.DialContext(ctx, cl.conn.addr.Network(), cl.conn.addr.String())
		if err != nil {
			cl.Logger.Warningf("retry connection %d", attempt)
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(cl.conn.delayFn(attempt)):
				continue
			}
		}

		return conn, nil
	}

	return nil, ErrMaxAttemptsExceeded
}

func (cl *Client) handleIncoming(ctx context.Context, conn net.Conn) error {
	buf := bufio.NewReaderSize(conn, 1024*1024) // 1MB buffer

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		frame, err := message.ReadFrame(buf, cl.conn.KeyData.Key)
		if err != nil && err == io.EOF {
			continue
		}
		if err != nil {
			return errors.Wrap(err, "read frame error")
		}
		err = cl.ProcessFrame(frame)
		if err != nil && err != context.Canceled {
			return errors.Wrap(err, "process frame error")
		}
	}

	//panic("unreachable code") // nolint: govet
}

func (cl *Client) handleOutgoing(ctx context.Context, conn net.Conn) (rerr error) {
	buffer := bufio.NewWriter(conn)
	defer func() {
		err := buffer.Flush()
		if err != nil {
			rerr = err
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case parcel := <-cl.outgoing:
			// TODO: Configurable write deadline
			// TODO: Think how write deadline will work with buffer
			err := conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
			if err != nil {
				return err
			}
			err = parcel.msg.Frame(cl.conn.KeyData).Send(buffer)
			if err != nil && isClosedConnError(err) {
				err = errConnClosed
			}
			parcel.done(ctx, err)
			if err != nil {
				return errors.Wrap(err, "failed to send frame")
			}
			if len(cl.outgoing) == 0 {
				_ = buffer.Flush() // nolint: errcheck
			}
		}
	}

	//panic("unreachable code") // nolint: govet
}

func (cl *Client) signKey(ctx context.Context) error {
	// Create key as a random string of 8 characters
	var key = []byte(crypt.RandString(8))

	// DesEncrypt symmetric key with server public key
	encryptedKey, err := crypt.EncryptDataWithServerPubKey(key, cl.crypt.Keys[sirenaXML.ServerPublicKey])
	if err != nil {
		return errors.Wrap(err, "encrypting data with server pubKey error")
	}
	msg, err := cl.newMessage(encryptedKey, false)
	if err != nil {
		return err
	}
	cl.conn.KeyData.Key = key
	return cl.send(ctx, msg)
}

func (cl *Client) SendMsg(ctx context.Context, data []byte) ([]byte, error) {
	msg, err := cl.newMessage(data, cl.conn.useZip)
	if err != nil {
		return nil, err
	}

	if err := cl.send(ctx, msg); err != nil {
		return nil, err
	}

	receivedMessage, err := cl.getResponseFromMsgPool(msg.MessageID)
	if err != nil {
		return nil, err
	}

	return receivedMessage.Payload, nil
}

func (cl *Client) GetKeyData() *message.KeyData {
	return cl.conn.KeyData
}

func isClosedConnError(err error) bool {
	// See: https://github.com/golang/go/issues/4373
	return err != nil && strings.Contains(err.Error(), "use of closed network connection")
}

func (cl *Client) getResponseFromMsgPool(msgID uint32) (*message.ReceivedMessage, error) {
	ctx, _ := context.WithTimeout(context.Background(), 40*time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case resp := <-cl.incoming[msgID]:
			cl.msgPool <- msgID
			return resp, nil
		}
	}
}

func (cl *Client) send(ctx context.Context, msg message.Message) error {
	sending := sending{
		errCh: make(chan error, 1),
		msg:   msg,
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case cl.outgoing <- sending:
	}

	return sending.wait(ctx)
}

func (cl *Client) newMessage(data []byte, zipIt bool) (*message.OutgoingMessage, error) {
	var (
		msgSign []byte
		err     error
	)

	if cl.conn.KeyData.Key == nil {
		msgSign, err = crypt.GeneratePrivateKeySignature(data, cl.crypt.Keys[sirenaXML.ClientPrivateKey], cl.crypt.ClientPrivKeyPass)
		if err != nil {
			return nil, err
		}
	}

	om := &message.OutgoingMessage{
		ClientID:         cl.conn.clientID,
		Message:          data,
		MessageSignature: msgSign,
		ZipIt:            zipIt,
	}

	if cl.conn.KeyData.Key != nil {
		om.MessageID = <-cl.msgPool
	}

	return om, nil
}

func (cl *Client) ProcessFrame(frame *message.Frame) error {
	msg := &message.ReceivedMessage{}
	msg.Decode(frame)

	switch msg.Type {
	case message.TypeSign:
		var err error
		keyData := &message.KeyData{}
		// DesDecrypt response
		keyData.Key, err = crypt.DecryptDataWithClientPrivateKey(msg.Payload[4:132], cl.crypt.Keys[sirenaXML.ClientPrivateKey], cl.crypt.ClientPrivKeyPass)
		if err != nil {
			return errors.Wrap(err, "decrypting data with client private key error")
		}
		keyData.ID = msg.KeyData.ID

		// Make sure request symmetric key = response symmatric key
		if string(keyData.Key) != string(cl.conn.KeyData.Key) {
			return errors.Errorf("Request symmetric key (%s) != response symmetric key(%s)", keyData.Key, cl.conn.KeyData.Key)
		}

		cl.conn.sessNum++
		cl.conn.initTime = time.Now()
		cl.conn.KeyData = keyData
		cl.Logger.Infof("started connection session number %d", cl.conn.sessNum)
	case message.TypeResponse:
		cl.incoming[msg.MessageID] <- msg
	case message.TypeError:
		return msg.Error()
	}

	return nil
}

func (cl *Client) dial(ctx context.Context) error {
	if cl.conn.dialer == nil {
		cl.conn.dialer = &net.Dialer{
			KeepAlive: 90 * time.Minute,
			Deadline:  time.Now().Add(90 * time.Minute),
		}
	}
	// Reconnect every half an hour to update symmetric key
	// https://wiki.sirena-travel.ru/xmlgate:01protocol:01encryption#sym_key_id
	deadline := time.Now().Add(90 * time.Minute)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	if cl.conn.reconnectAttempts == 0 {
		cl.conn.reconnectAttempts = math.MaxUint32
	}

	if cl.conn.delayFn == nil {
		initialDelay, maxDelay := 1*time.Second, 5*time.Second
		cl.conn.delayFn = LinearDelay(initialDelay, maxDelay)
	}

	conn, err := cl.establishConn(ctx)
	switch err {
	case ErrMaxAttemptsExceeded, context.Canceled, context.DeadlineExceeded:
		return err
	default:
		if err != nil {
			panic(fmt.Sprintf("unreachable code: %v", err))
		}
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		g, gctx := errgroup.WithContext(ctx)

		g.Go(func() error {
			return cl.handleIncoming(gctx, conn)
		})

		g.Go(func() error {
			handleErr := cl.handleOutgoing(gctx, conn)

			// Close connection only after all outgoing messages have been sent.
			// Connection closing here also used to interrupt reading from conn inside handleIncoming,
			// so it can't be placed after Wait().
			//closeErr := conn.Close()
			//if closeErr != nil {
			//	//return closeErr if it has more information about problems
			//	return closeErr
			//}

			return handleErr
		})

		gerr := g.Wait()
		if gerr != nil &&
			gerr != context.Canceled &&
			gerr != context.DeadlineExceeded &&
			gerr != io.EOF &&
			!isClosedConnError(gerr) {
			// TODO: Notify gerr
			// fmt.Println("TODO: Notify gerr", gerr)
			cl.Logger.Error(gerr)
		}

		// In any case, the first reconnection attempt will be made immediately.
	}

	//panic("unreachable code") // nolint: govet
}

func (cl *Client) WaitKeySign(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if cl.conn.KeyData.ID != 0 {
			break
		}
	}
	return nil
}
