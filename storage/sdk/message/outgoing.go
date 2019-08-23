package message

// OutgoingMessage represents application level data.
type OutgoingMessage struct {
	MessageID        uint32
	ClientID         uint16
	KeyData          KeyData
	flags            *HeaderFlags
	Message          []byte
	MessageSignature []byte
	ZipIt            bool
}

func (o *OutgoingMessage) Decode(f *Frame) {
	return nil
}

func (o *OutgoingMessage) Frame(kd *KeyData) *Frame {
	header := makeHeader(o, kd)
	var data []byte
	data = append(data, header.ToBytes()...)
	data = append(data, o.Message...)

	if !mustBeSymmetricKeyEncrypted(kd) {
		data = append(data, encryptedMsgLengthToBytes(o.Message)...)
		data = append(data, o.MessageSignature...)
	}

	return &Frame{Payload: data}
}

func mustBeSymmetricKeyEncrypted(kd *KeyData) bool {
	return kd.ID != 0 && kd.Key != nil
}
