package sdk

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *storage) GetKeyInfo(req []byte) (*structs.KeyInfoResponse, error) {
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
