package message

import (
	"github.com/pkg/errors"
)

// ReceivedMessage represents application level data.
type ReceivedMessage struct {
	Type      MsgType
	MessageID uint32
	ClientID  uint16
	KeyData   *KeyData
	Payload   []byte
	flags     *HeaderFlags
	Err       *string
}

type MsgType int

const (
	TypeError MsgType = iota
	TypeSign
	TypeResponse
)

type KeyData struct {
	ID  uint32 `json:"id"`
	Key []byte `json:"key"`
}

func (r *ReceivedMessage) Frame(kd *KeyData) *Frame {
	panic("fuck!")
	//return &Frame{
	//	Header: &Header{
	//		Flags:         r.flags,
	//		MessageID:     r.MessageID,
	//		MessageLength: uint32(len(r.Payload)),
	//	},
	//	Payload: r.Payload,
	//}
}

func (r *ReceivedMessage) Decode(f *Frame) {
	r.MessageID = f.Header.MessageID
	r.ClientID = f.Header.ClientID
	r.Payload = f.Payload
	r.flags = f.Header.Flags
	r.KeyData = &KeyData{ID: f.Header.KeyID}
	r.setType(f)
}

func (r *ReceivedMessage) setType(f *Frame) {
	r.Type = TypeResponse

	if f.Header.MessageID == 0 {
		r.Type = TypeSign
	}

	if f.Header.RequestNoHandled {
		r.Type = TypeError
	}
}

func (r *ReceivedMessage) Error() error {
	return errors.New("receive message error: " + *r.Err)
}
