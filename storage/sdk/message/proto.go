package message

// Payload is the interface implemented by an object that can decode and encode
// a particular message.
type Message interface {
	Decode(f *Frame)
	Frame(kd *KeyData) *Frame
}
