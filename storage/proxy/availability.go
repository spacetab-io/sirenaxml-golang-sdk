package proxy

import (
	"encoding/xml"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *storage) GetAvailability(req []byte) (*structs.AvailabilityResponse, error) {
	resp, err := s.sendMsg(req)
	if err != nil {
		return nil, err
	}

	var availability structs.AvailabilityResponse
	err = xml.Unmarshal(resp, &availability)

	return &availability, err
}
