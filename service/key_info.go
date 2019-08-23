package service

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *service) KeyInfo() (*structs.KeyManager, *structs.Error, error) {
	reqXML, err := xml.Marshal(&structs.KeyInfoRequest{})
	if err != nil {
		return nil, nil, err
	}
	keyAnswer, err := s.sdk.GetCurrentKeyInfo(reqXML)
	if err != nil {
		return nil, nil, err
	}

	return &keyAnswer.Answer.KeyInfo.KeyManager, nil, nil
}
