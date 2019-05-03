package client

import "github.com/tmconsulting/sirenaxml-golang-sdk/logs"

type HeaderFlags struct {
	val byte
}

func NewHeaderFlags(preset byte) *HeaderFlags {
	hf := &HeaderFlags{val: preset}
	logs.Log.Debugf("reading response header flags: ZippedRequest: %v, ZippedResponse: %v, EncryptSymmetric: %v, EncryptPublic: %v", hf.Has(ZippedRequest), hf.Has(ZippedResponse), hf.Has(EncryptSymmetric), hf.Has(EncryptPublic))
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
