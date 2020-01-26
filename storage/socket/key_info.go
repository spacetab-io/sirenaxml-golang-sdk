package socket

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/socket/client"
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

func (s *storage) GetKeyData() (*client.KeyData, error) {
	kd := s.c.GetKeyData()
	return &kd, nil
}
