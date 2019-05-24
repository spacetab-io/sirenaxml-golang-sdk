package client

import (
	"bufio"

	"github.com/pkg/errors"
)

func readPacket(reader *bufio.Reader) error {
	responseHeaderBytes := make([]byte, 100)
	if _, err := reader.Read(responseHeaderBytes); err != nil {
		if err.Error() == "EOF" {
			return errors.New("your ip is not in white list")
		} else {
			return errors.Wrap(err, "receiving header error")
		}
	}

	header, err := parseHeader(responseHeaderBytes)
	if err != nil {
		return err
	}

	responseMessageBytes := make([]byte, header.MessageLength)
	_, err = reader.Read(responseMessageBytes)
	if err != nil {
		return err
	}

	return respPool.SavePacket(header.MessageID, &Packet{
		header:  header,
		message: responseMessageBytes,
	})
}
