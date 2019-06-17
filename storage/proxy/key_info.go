package proxy

import (
	"encoding/json"
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/client"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *storage) GetKeyData() (*client.KeyData, error) {
	resp, err := s.r.R().Get(s.proxyPath + "/key_info")
	if err != nil || resp.StatusCode() != 200 {
		if err == nil {
			return nil, nil
		}
		return nil, nil
	}
	var keyData client.KeyData
	err = json.Unmarshal(resp.Body(), &keyData)

	return &keyData, err
}

func (s *storage) GetCurrentKeyInfo(req []byte) (*structs.KeyInfoResponse, error) {
	resp, err := s.sendMsg(req)
	if err != nil {
		return nil, err
	}
	var keyInfo structs.KeyInfoResponse
	err = xml.Unmarshal(resp, &keyInfo)

	return &keyInfo, err
}
