package client

import (
	"net"

	sirenaXML "github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
)

// Option describes a functional option for configuring the Client.
type Option func(*Client)

// Dialer sets the Dialer of Client.
func Dialer(dialer CustomDialer) Option {
	return func(cl *Client) {
		cl.conn.dialer = dialer
	}
}

func SetUseZip(useZip bool) Option {
	return func(cl *Client) {
		cl.conn.useZip = useZip
	}
}

func SetClientID(clientID uint16) Option {
	return func(cl *Client) {
		cl.conn.clientID = clientID
	}
}

func SetLogger(l logs.LogWriter) Option {
	return func(cl *Client) {
		cl.Logger = l
	}
}

func SetKeys(kd *sirenaXML.KeysData) Option {
	return func(cl *Client) {
		cl.crypt = kd
	}
}
func SetAddr(addr *net.TCPAddr) Option {
	return func(cl *Client) {
		cl.conn.addr = addr
	}
}

func SetMaxConnections(mc uint32) Option {
	return func(cl *Client) {
		cl.conn.maxConnections = mc
	}
}
