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
	responseHeaderBytes := make([]byte, 100)
	if _, err := reader.Read(responseHeaderBytes); err != nil {
		return err
	}

	header, err := parseHeader(responseHeaderBytes)
	if err != nil {
		return err
	}

	message, err := readMessage(header, reader, c.socket.KeyData.Key)
	if err != nil {
		// @TODO send error packet
		return err
	}

	if err := respPool.SavePacket(header.MessageID, &Packet{header: header, message: message}); err != nil {
		// @TODO send error packet
		return err
	}

	return nil
}

func readMessage(header *Header, reader *bufio.Reader, key []byte) ([]byte, error) {
	responseMessageBytes := make([]byte, header.MessageLength)
	_, err := io.ReadFull(reader, responseMessageBytes)
	if err != nil {
		return nil, err
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
		responseMessageBytes, err = crypt.DesDecrypt(responseMessageBytes, key)
		if err != nil {
			return nil, err
		}
	}

	return responseMessageBytes, nil
}
