package client

import (
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
)

type Header struct {
	CreatedAt        uint32
	MessageID        uint32
	MessageLength    uint32
	ClientID         uint16
	KeyID            uint32
	Flags            *HeaderFlags
	RequestNoHandled bool
}

// HeaderOffsets holds information about header offset lengths
var HeaderOffsets = map[int]int{
	0: 0,
	1: 4,
	2: 8,
	3: 12,
	4: 44,
	5: 46,
	6: 47,
	7: 48,
	8: 52,
}

// ParseHeader parses bytes into header
func parseHeader(data []byte) (*Header, error) {
	header := &Header{
		MessageLength:    binary.BigEndian.Uint32(data[HeaderOffsets[0]:]),
		CreatedAt:        binary.BigEndian.Uint32(data[HeaderOffsets[1]:]),
		MessageID:        binary.BigEndian.Uint32(data[HeaderOffsets[2]:]),
		ClientID:         binary.BigEndian.Uint16(data[HeaderOffsets[4]:]),
		Flags:            NewHeaderFlags(data[HeaderOffsets[5]]),
		RequestNoHandled: 0x01&data[HeaderOffsets[6]] != 0,
		KeyID:            binary.BigEndian.Uint32(data[HeaderOffsets[7]:]),
	}

	if header.MessageID == 0 {
		return nil, errors.New("messageID is not set")
	}

	if header.MessageLength == 0 {
		return nil, errors.New(fmt.Sprintf("[%d] empty message", header.MessageID))
	}

	if header.RequestNoHandled {
		return nil, errors.New(fmt.Sprintf("[%d] request not handled", header.MessageID))
	}

	return header, nil
}

// ToBytes converts header into bytes
func (h *Header) ToBytes() []byte {
	headerBytes := make([]byte, 100)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[0]:], h.MessageLength)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[1]:], h.CreatedAt)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[2]:], h.MessageID)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[3]:], 0)
	binary.BigEndian.PutUint16(headerBytes[HeaderOffsets[4]:], h.ClientID)
	headerBytes[HeaderOffsets[5]] = h.Flags.ToByte()
	headerBytes[HeaderOffsets[6]] = 0
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[7]:], h.KeyID)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[8]:], 0)
	return headerBytes
}

func (h *Header) setFlags(cfg *sirenaXML.Config, sign bool) {
	h.Flags = &HeaderFlags{}

	// it will be easier to manage zipped status of request and response in one config attribute
	if cfg.ZippedMessaging {
		h.Flags.Set(ZippedRequest)
		h.Flags.Set(ZippedResponse)
	}

	if sign {
		h.MessageLength += 4 + 128
		h.Flags.Set(EncryptPublic)
	} else {
		h.Flags.Set(EncryptSymmetric)
	}
}
