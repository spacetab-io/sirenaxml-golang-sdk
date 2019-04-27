package sirena

import (
	"encoding/binary"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

const (
	// HeaderSize is a header size
	HeaderSize int32 = 100
)

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

// Header is a header in Sirena request
type Header struct {
	CreatedAt        uint32
	MessageID        uint32
	MessageLength    uint32
	ClientID         uint16
	KeyID            uint32
	Flags            *HeaderFlags
	RequestNoHandled bool
}

// NewHeaderParams holds params for new header
type NewHeaderParams struct {
	ClientID            uint16
	KeyID               uint32
	MessageLength       uint32
	MessageID           uint32
	MessageIsZipped     bool
	ResponseCanBeZipped bool
	UseEncrypt          bool
	UsePublic           bool
	UseSymmetric        bool
}

// NewHeader creates new header for provided message
func NewHeader(params *NewHeaderParams) *Header {
	return &Header{
		MessageLength: params.MessageLength,
		CreatedAt:     uint32(time.Now().Unix()),
		ClientID:      params.ClientID,
		MessageID:     uint32(params.MessageID),
		Flags:         setFlags(params),
		KeyID:         params.KeyID,
	}
}

func setFlags(params *NewHeaderParams) *HeaderFlags {
	flags := &HeaderFlags{}

	// flags = 0x22
	if params.MessageIsZipped {
		flags.Set(ZippedRequest)
	}

	if params.ResponseCanBeZipped {
		flags.Set(ZippedResponse)
	}
	if params.UseSymmetric {
		flags.Set(EncryptSymmetric)
	}
	if params.UseEncrypt {
		params.MessageLength += 4 + 128
		flags.Set(EncryptPublic)
	}
	if params.UsePublic {
		flags.Set(EncryptPublic)
	}
	return flags
}

// ToBytes converts header into bytes
func (h *Header) ToBytes() []byte {
	headerBytes := make([]byte, HeaderSize)
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

// ParseHeader parses bytes into header
func ParseHeader(data []byte) (*Header, error) {
	rh := &Header{
		MessageLength:    binary.BigEndian.Uint32(data[HeaderOffsets[0]:]),
		CreatedAt:        binary.BigEndian.Uint32(data[HeaderOffsets[1]:]),
		MessageID:        binary.BigEndian.Uint32(data[HeaderOffsets[2]:]),
		ClientID:         binary.BigEndian.Uint16(data[HeaderOffsets[4]:]),
		Flags:            NewHeaderFlags(data[HeaderOffsets[5]]),
		RequestNoHandled: 0x01&data[HeaderOffsets[6]] != 0,
		KeyID:            binary.BigEndian.Uint32(data[HeaderOffsets[7]:]),
	}

	if rh.MessageLength == 0 {
		return nil, errors.Errorf("sirena response header doesn't include message length: %s", spew.Sdump(rh))
	}

	if rh.RequestNoHandled {
		return nil, errors.New("sirena response not handled for some reason")
	}

	return rh, nil
}

// MakeSubHeader returns sub header holding length of data passed
func MakeSubHeader(data []byte) []byte {
	subHeader := make([]byte, 4)
	binary.BigEndian.PutUint32(subHeader[0:], uint32(len(data)))
	return subHeader
}
