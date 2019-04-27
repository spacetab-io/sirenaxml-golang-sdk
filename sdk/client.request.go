package sdk

import (
	"bytes"
	"compress/zlib"
	"math/rand"
	"time"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
	"github.com/tmconsulting/sirenaxml-golang-sdk/des"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
)

func (client *SirenaClient) NewSignRequest(key []byte) (*Request, error) {
	logs.Log.Debugf("Trying to sign DES key %s with Sirena", key)

	// Encrypt symmetric key with server public key
	encryptedKey, err := crypt.EncryptDataWithServerPubKey(key, client.Config.ServerPublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "encrypting data with server pubKey error")
	}

	h := NewHeader(&NewHeaderParams{
		ClientID:      client.Config.ClientID,
		MessageLength: uint32(len(encryptedKey)),
		UseEncrypt:    true,
	})
	sh := MakeSubHeader(encryptedKey)
	ms, err := crypt.GeneratePrivateKeySignature(encryptedKey, client.Config.ClientPrivateKey, client.Config.ClientPrivateKeyPassword)
	if err != nil {
		return nil, err
	}
	return &Request{
		Header:           h,
		SubHeader:        sh,
		Message:          encryptedKey,
		MessageSignature: ms,
	}, nil
}

func (client *SirenaClient) NewRequest(msg []byte) (*Request, error) {

	if client.Config.ZipRequests {
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

	rand.Seed(time.Now().Unix())
	logs.Log.Debugf("encrypt message with symmetric DES key %s (keyID %d)", client.Key, client.KeyID)
	msg, err := des.Encrypt(msg, client.Key)
	if err != nil {
		return nil, err
	}

	return &Request{
		Header: NewHeader(&NewHeaderParams{
			KeyID:               client.KeyID,
			ClientID:            client.Config.ClientID,
			MessageLength:       uint32(len(msg)),
			UseSymmetric:        true,
			MessageIsZipped:     client.Config.ZipRequests,
			ResponseCanBeZipped: client.Config.ZipResponses,
		}),
		Message: msg,
	}, nil
}
