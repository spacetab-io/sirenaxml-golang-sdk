package client

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
)

func (c *Channel) readPacket(reader *bufio.Reader) error {
	if reader == nil {
		return errors.New("reader is closed")
	}
	var (
		err     error
		header  *Header
		message []byte
	)
	responseHeaderBytes := make([]byte, 100)
	if _, err = io.ReadFull(reader, responseHeaderBytes); err != nil {
		return errors.Wrap(err, "read header error")
	}

	if header, err = parseHeader(responseHeaderBytes); err != nil {
		return errors.Wrap(err, "parse header error")
	}

	if message, err = readMessage(header, reader, c.socket.KeyData.Key); err != nil {
		return errors.Wrap(err, "read message error")
	}

	if err := respPool.SavePacket(header.MessageID, &Packet{header: header, message: message}); err != nil {
		return errors.Wrap(err, "save packet error")
	}

	return nil
}

func readMessage(header *Header, reader *bufio.Reader, key []byte) ([]byte, error) {
	var (
		err                  error
		responseMessageBytes = make([]byte, header.MessageLength)
	)
	if _, err = io.ReadFull(reader, responseMessageBytes); err != nil {
		return nil, errors.Wrap(err, "read response message bytes error")
	}

	// unzip message if it was zipped
	if header.Flags.Has(ZippedResponse) {
		b := bytes.NewReader(responseMessageBytes)
		z, err := zlib.NewReader(b)
		if err != nil {
			return nil, errors.Wrap(err, "zlib new reader error")
		}
		responseMessageBytes, err = ioutil.ReadAll(z)
		if err != nil {
			return nil, err
		}
		err = z.Close()
		if err != nil {
			return nil, err
		}
	}

	// decrypt symmetric key encrypted message
	if key != nil && header.Flags.Has(EncryptSymmetric) {
		if responseMessageBytes, err = crypt.DesDecrypt(responseMessageBytes, key); err != nil {
			return nil, err
		}
	}

	return responseMessageBytes, nil
}
