package proxy

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/errors"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *storage) GetAvailability(req []byte) (*structs.AvailabilityResponse, *structs.Error, error) {
	response, respError, err := s.sendMsg(req)
	if hasErr, respErr, err := errors.CheckErrors(respError, err); hasErr {
		return nil, respErr, err
	}

	var availability structs.AvailabilityResponse
	err = xml.Unmarshal(response, &availability)

	return &availability, nil, err
}
