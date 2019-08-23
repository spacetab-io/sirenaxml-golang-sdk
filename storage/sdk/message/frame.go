package message

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"io"
	"io/ioutil"
	"net"
	"syscall"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
)

type Frame struct {
	Header *Header
	//subHeader        []byte
	//Payload          []byte
	//msgSignature []byte
	Payload []byte
}

func (f *Frame) Send(w io.Writer) error {

	if _, err := w.Write(f.Payload); err != nil {
		return errors.Wrap(err, "packet write error")
	}

	return nil
}

func ReadFrame(reader *bufio.Reader, key []byte) (*Frame, error) {
	var (
		err   error
		frame = &Frame{}
	)
	frame.Header, err = readHeader(reader)
	if err != nil && err == io.EOF {
		return nil, err
	}
	if err != nil {
		return nil, errors.Wrap(err, "read header error")
	}

	frame.Payload, err = readMessage(frame.Header, reader, key)
	if err != nil {
		return nil, errors.Wrap(err, "read message error")
	}

	return frame, nil
}

func readHeader(reader *bufio.Reader) (*Header, error) {
	responseHeaderBytes := make([]byte, 100)
	n, err := io.ReadFull(reader, responseHeaderBytes)
	if err != nil && err == io.EOF {
		return nil, err
	}
	if err != nil {
		return nil, errors.Wrap(err, "read from buffer error")
	}

	if n == 0 {
		return nil, errors.New("no data reads")
	}

	header, err := parseHeader(responseHeaderBytes)
	if err != nil {
		return nil, errors.Wrap(err, "parse header error")
	}

	return header, nil
}

func econResetErr(err error) bool {
	switch t := err.(type) {
	case *net.OpError:
		if t.Op == "read" {
			return true
		}

	case syscall.Errno:
		if t == syscall.ECONNRESET {
			return true
		}
	}

	return false
}

func readMessage(header *Header, reader *bufio.Reader, key []byte) ([]byte, error) {
	var (
		err                  error
		responseMessageBytes = make([]byte, header.MessageLength)
	)
	n, err := io.ReadFull(reader, responseMessageBytes)
	if err != nil {
		return nil, errors.Wrap(err, "read response message bytes error")
	}
	if n == 0 {
		return nil, errors.Wrap(err, "no data in response")
	}

	// unzip message if it was Zipped
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
