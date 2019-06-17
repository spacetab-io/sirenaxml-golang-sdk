package service

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *service) KeyInfo() (*structs.KeyInfoResponse, error) {
	reqXML, err := xml.Marshal(&structs.KeyInfoRequest{})
	if err != nil {
		return nil, err
	}
	return s.sdk.GetCurrentKeyInfo(reqXML)
}
