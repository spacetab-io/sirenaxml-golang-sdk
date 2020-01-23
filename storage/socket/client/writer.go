package client

import (
	"bufio"
	"github.com/pkg/errors"
)

func writePacket(writer *bufio.Writer, packet *Packet) error {
	var data []byte
	data = append(data, packet.header.ToBytes()...)
	if len(packet.subHeader) > 0 {
		data = append(data, packet.subHeader...)
	}
	data = append(data, packet.message...)
	if len(packet.messageSignature) > 0 {
		data = append(data, packet.messageSignature...)
	}
	if _, err := writer.Write(data); err != nil {
		return errors.Wrap(err, "packet write error")
	}

	return nil
}
