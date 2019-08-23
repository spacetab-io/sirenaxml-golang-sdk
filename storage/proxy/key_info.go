package proxy

import (
	"encoding/json"
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/errors"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/message"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *storage) GetKeyData() (*message.KeyData, error) {
	resp, err := s.r.R().Get(s.proxyPath + "/key_info")
	if err != nil || resp.StatusCode() != 200 {
		if err == nil {
			return nil, nil
		}
		return nil, nil
	}
	var keyData message.KeyData
	err = json.Unmarshal(resp.Body(), &keyData)

	return &keyData, err
}

func (s *storage) GetCurrentKeyInfo(req []byte) (*structs.KeyInfoResponse, *structs.Error, error) {
	resp, respError, err := s.sendMsg(req)
	if hasErr, respErr, err := errors.CheckErrors(respError, err); hasErr {
		return nil, respErr, err
	}

	var keyInfo structs.KeyInfoResponse
	err = xml.Unmarshal(resp, &keyInfo)

	return &keyInfo, nil, err
}
