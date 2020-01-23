package socket

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *storage) GetAvailability(req []byte) (*structs.AvailabilityResponse, error) {
	// Create Sirena request
	response, err := s.SendRawRequest(req)
	if err != nil {
		return nil, err
	}

	var availabilityResponse structs.AvailabilityResponse
	if err := xml.Unmarshal(response, &availabilityResponse); err != nil {
		return nil, err
	}

	return &availabilityResponse, nil
}
