package sdk

import (
	"encoding/xml"

	errs "github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/message"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *storage) GetCurrentKeyInfo(req []byte) (*structs.KeyInfoResponse, error) {
	// Create Sirena request
	response, err := s.SendRawRequest(req)
	if err != nil {
		return nil, err
	}

	var keyInfo structs.KeyInfoResponse
	if err := xml.Unmarshal(response, &keyInfo); err != nil {
		return nil, err
	}

	return &keyInfo, nil
}

func (s *storage) GetKeyData() (*message.KeyData, error) {
	kd := s.c.GetKeyData()
	if &kd == nil {
		return nil, errs.New("no key data occurred")
	}

	return kd, nil
}
