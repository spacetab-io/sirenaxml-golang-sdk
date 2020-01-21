package client

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"time"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
)

func (c *Channel) signKey() error {
	// Create key as a random string of 8 characters
	var key = []byte(crypt.RandString(8))

	// Create Sirena request
	request, err := c.newSignRequestPacket(key)
	if err != nil {
		return errors.Wrap(err, "making request error")
	}

	// Send request to Sirena
	c.sendPacket(request)

	// oneshot receiving action
	if err := c.readPacket(bufio.NewReader(c.socket.conn)); err != nil {
		return errors.Wrap(err, "receive key error")
	}

	response := getResponseFromMsgPool(request.header.MessageID)

	// DesDecrypt response
	c.socket.KeyData.Key, err = crypt.DecryptDataWithClientPrivateKey(response.message[4:132], []byte(c.cfg.ClientPrivateKey), c.cfg.ClientPrivateKeyPassword)
	if err != nil {
		return errors.Wrap(err, "decrypting data with client private key error")
	}
	c.socket.KeyData.ID = response.header.KeyID

	// Make sure request symmetric key = response symmatric key
	if string(key) != string(c.socket.KeyData.Key) {
		return errors.Errorf("Request symmetric key (%s) != response symmetric key(%s)", key, c.socket.KeyData.Key)
	}

	return nil
}

func (c *Channel) newSignRequestPacket(key []byte) (*Packet, error) {
	// DesEncrypt symmetric key with server public key
	encryptedKey, err := crypt.EncryptDataWithServerPubKey(key, []byte(c.cfg.ServerPublicKey))
	if err != nil {
		return nil, errors.Wrap(err, "encrypting data with server pubKey error")
	}

	initCfg := c.cfg
	initCfg.ZippedMessaging = false

	return NewPacket(c, encryptedKey, c.socket.KeyData.ID)
}

func NewPacket(c *Channel, key []byte, keyID uint32) (*Packet, error) {
	var err error
	p := &Packet{}
	p.makeHeader(c, key, keyID)
	p.messageSignature, err = crypt.GeneratePrivateKeySignature(key, []byte(c.cfg.ClientPrivateKey), c.cfg.ClientPrivateKeyPassword)
	if err != nil {
		return nil, err
	}
	p.message = key
	return p, err
}

func (c *Channel) NewRequest(msg []byte) (*Packet, error) {
	var err error
	if c.cfg.ZippedMessaging {
		buf := new(bytes.Buffer)
		w := zlib.NewWriter(buf)
		_, err = w.Write(msg)
		if err != nil {
			return nil, err
		}
		err = w.Close()
		if err != nil {
			return nil, err
		}
		msg = buf.Bytes()
	}

	if c.socket.KeyData.Key != nil {
		msg, err = crypt.DesEncrypt(msg, c.socket.KeyData.Key)
		if err != nil {
			return nil, err
		}
	}

	p := &Packet{}
	p.makeHeader(c, msg, c.socket.KeyData.ID)
	p.message = msg
	return p, err
}

func (p *Packet) makeHeader(c *Channel, key []byte, keyID uint32) {
	msgID := msgPool.GetMsgID()
	sign := false
	p.header = &Header{
		MessageID:     msgID,
		ClientID:      c.cfg.ClientID,
		MessageLength: uint32(len(key)),
		CreatedAt:     uint32(time.Now().Unix()),
	}
	if keyID == 0 {
		sign = true
		p.makeSubHeader(key)
	} else {
		p.header.KeyID = keyID
	}

	p.header.setFlags(c, sign)
}

func (p *Packet) makeSubHeader(data []byte) {
	p.subHeader = make([]byte, 4)
	binary.BigEndian.PutUint32(p.subHeader[0:], uint32(len(data)))
}
