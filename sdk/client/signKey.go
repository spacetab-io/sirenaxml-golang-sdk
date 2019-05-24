package client

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"time"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
	"github.com/tmconsulting/sirenaxml-golang-sdk/des"
	"github.com/tmconsulting/sirenaxml-golang-sdk/random"
)

func (c *Channel) signKey() error {
	// Create key as a random string of 8 characters
	var key = []byte(random.String(8))

	// Create Sirena request
	request, err := c.newSignRequestPacket(key)
	if err != nil {
		return errors.Wrap(err, "making request error")
	}

	// Send request to Sirena
	c.sendPacket(request)

	// oneshot receiving action
	if err := receive(c); err != nil {
		return errors.Wrap(err, "schedule receive error")
	}

	response := getResponseFromMsgPool(request.header.MessageID)

	// Decrypt response
	c.Key, err = crypt.DecryptDataWithClientPrivateKey(response.message[4:132], c.cfg.ClientPrivateKey, c.cfg.ClientPrivateKeyPassword)
	if err != nil {
		return errors.Wrap(err, "decrypting data with client private key error")
	}
	c.KeyID = response.header.KeyID

	// Make sure request symmetric key = response symmatric key
	if string(key) != string(c.Key) {
		return errors.Errorf("Request symmetric key (%s) != response symmetric key(%s)", key, c.Key)
	}

	c.cfg.UseSymmetricKeyCrypt = true
	return nil
}

func (c *Channel) newSignRequestPacket(key []byte) (*Packet, error) {
	// Encrypt symmetric key with server public key
	encryptedKey, err := crypt.EncryptDataWithServerPubKey(key, c.cfg.ServerPublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "encrypting data with server pubKey error")
	}

	initCfg := c.cfg
	initCfg.UseSymmetricKeyCrypt = false
	initCfg.UsePublicKeyCrypt = false
	initCfg.ZipResponses = false
	initCfg.ZipRequests = false

	return NewPacket(initCfg, encryptedKey)
}

func NewPacket(cfg *configuration.SirenaConfig, key []byte) (*Packet, error) {
	var err error
	p := &Packet{}
	p.makeHeader(cfg, key)
	p.messageSignature, err = crypt.GeneratePrivateKeySignature(key, cfg.ClientPrivateKey, cfg.ClientPrivateKeyPassword)
	if err != nil {
		return nil, err
	}
	p.message = key
	return p, err
}

func (c *Channel) NewRequest(msg []byte) (*Packet, error) {

	if c.cfg.ZipRequests {
		buf := new(bytes.Buffer)
		w := zlib.NewWriter(buf)
		_, err := w.Write(msg)
		if err != nil {
			return nil, err
		}
		err = w.Close()
		if err != nil {
			return nil, err
		}
		msg = buf.Bytes()
	}

	//logs.Log.Debugf("encrypt message with symmetric DES key %s (keyID %d)", client.Key, client.KeyID)
	msg, err := des.Encrypt(msg, c.Key)
	if err != nil {
		return nil, err
	}

	p := &Packet{}
	p.makeMsgHeader(c.cfg, msg, c.KeyID)
	//logs.Log.Debug("GeneratePrivateKeySignature")
	p.message = msg
	return p, err
}

func (p *Packet) makeHeader(cfg *configuration.SirenaConfig, key []byte) {
	msgID := msgPool.GetMsgID()
	p.header = &Header{
		MessageID:     msgID,
		ClientID:      cfg.ClientID,
		MessageLength: uint32(len(key)),
		CreatedAt:     uint32(time.Now().Unix()),
	}
	p.header.setFlags(cfg, true)
	p.makeSubHeader(key)
}
func (p *Packet) makeMsgHeader(cfg *configuration.SirenaConfig, key []byte, keyID uint32) {
	msgID := msgPool.GetMsgID()
	p.header = &Header{
		MessageID:     msgID,
		ClientID:      cfg.ClientID,
		MessageLength: uint32(len(key)),
		CreatedAt:     uint32(time.Now().Unix()),
		KeyID:         keyID,
	}
	p.header.setFlags(cfg, true)
	p.makeSubHeader(key)
}

func (p *Packet) makeSubHeader(data []byte) {
	p.subHeader = make([]byte, 4)
	binary.BigEndian.PutUint32(p.subHeader[0:], uint32(len(data)))
}
