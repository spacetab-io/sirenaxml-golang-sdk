package sirena

import (
	"encoding/binary"
	"math/rand"
	"time"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logger"

	"github.com/tmconsulting/sirenaxml-golang-sdk/config"
	"github.com/tmconsulting/sirenaxml-golang-sdk/utils"
)

const (
	// HeaderSize is a header size
	HeaderSize int32 = 100
	// ResponseZipped is a flag saying response is zipped
	ResponseZipped byte = 0x10
	// RequestZipped is a flag saying request is zipped
	RequestZipped byte = 0x04
	// EncryptSymmetric is a flag saying message is encrypted by symmetric key (DES)
	EncryptSymmetric byte = 0x08
	// EncryptPublic is a flag saying message is encrypted by public key (RSA)
	EncryptPublic byte = 0x40
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
	CreatedAt     uint32
	MessageID     uint32
	MessageLength uint32
	ClientID      uint16
	KeyID         uint32
	Flags         byte
}

// NewHeaderParams holds params for new header
type NewHeaderParams struct {
	Message      []byte
	KeyID        uint32
	MessageID    uint32
	CanBeZipped  bool
	UseEncrypt   bool
	UsePublic    bool
	UseSymmetric bool
}

// NewHeader creates new header for provided message
func NewHeader(params NewHeaderParams) *Header {
	if len(params.Message) == 0 {
		return nil
	}

	config := config.Get()

	msgLength := uint32(len(params.Message))
	rand.Seed(time.Now().Unix())

	var flags byte

	// flags = 0x22
	if params.CanBeZipped {
		flags = RequestZipped
	}
	if params.UseEncrypt {
		msgLength += 4 + 128
		flags = EncryptPublic
	}
	if params.UsePublic {
		flags = EncryptPublic
	}
	if params.UseSymmetric {
		flags = EncryptSymmetric
	}

	logger := logger.Get()

	sirenaClientID, err := utils.String2Uint16(config.SirenaClientID)
	if err != nil {
		logger.Error(err)
		return nil
	}

	return &Header{
		MessageLength: msgLength,
		CreatedAt:     uint32(time.Now().Unix()),
		ClientID:      sirenaClientID,
		KeyID:         params.KeyID,
		MessageID:     uint32(params.MessageID),
		Flags:         flags,
	}
}

// ToBytes converts header into bytes
func (h *Header) ToBytes() []byte {
	headerBytes := make([]byte, HeaderSize)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[0]:], h.MessageLength)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[1]:], h.CreatedAt)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[2]:], h.MessageID)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[3]:], 0)
	binary.BigEndian.PutUint16(headerBytes[HeaderOffsets[4]:], h.ClientID)
	headerBytes[HeaderOffsets[5]] = h.Flags
	headerBytes[HeaderOffsets[6]] = 0
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[7]:], h.KeyID)
	binary.BigEndian.PutUint32(headerBytes[HeaderOffsets[8]:], 0)
	return headerBytes
}

// ParseHeader parses bytes into header
func ParseHeader(data []byte) Header {
	h := Header{}
	h.MessageLength = binary.BigEndian.Uint32(data[HeaderOffsets[0]:])
	h.CreatedAt = binary.BigEndian.Uint32(data[HeaderOffsets[1]:])
	h.MessageID = binary.BigEndian.Uint32(data[HeaderOffsets[2]:])
	h.ClientID = binary.BigEndian.Uint16(data[HeaderOffsets[4]:])
	h.Flags = data[HeaderOffsets[5]]
	h.KeyID = binary.BigEndian.Uint32(data[HeaderOffsets[7]:])
	return h
}

// MakeSubHeader returns sub header holding length of data passed
func MakeSubHeader(data []byte) []byte {
	subHeader := make([]byte, 4)
	binary.BigEndian.PutUint32(subHeader[0:], uint32(len(data)))
	return subHeader
}
