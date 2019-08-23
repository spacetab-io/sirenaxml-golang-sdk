package message

type HeaderFlags struct {
	val byte
}

const (
	ZippedRequest    byte = 0x04 // ZippedRequest is a flag saying request is gzipped
	ZippedResponse        = 0x10 // ZippedResponse is a flag saying response can be gzipped
	EncryptSymmetric      = 0x08 // EncryptSymmetric is a flag saying message is encrypted by symmetric key (DES)
	EncryptPublic         = 0x40 // EncryptPublic is a flag saying message is encrypted by public key (RSA)
)

func NewHeaderFlags(preset byte) *HeaderFlags {
	hf := &HeaderFlags{val: preset}
	return hf
}

func (hf *HeaderFlags) Set(flag byte) {
	hf.val |= flag
}

func (hf *HeaderFlags) Clear(flag byte) {
	hf.val = hf.val &^ flag
}

func (hf *HeaderFlags) Toggle(flag byte) {
	hf.val ^= flag
}

func (hf *HeaderFlags) Has(flag byte) bool {
	return hf.val&flag != 0
}

func (hf *HeaderFlags) ToByte() byte {
	return hf.val
}
