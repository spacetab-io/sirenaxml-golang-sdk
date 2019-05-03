package sdk

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *storage) GetAvailability(req []byte) (*structs.AvailabilityResponse, error) {
	// Create Sirena request
	request, err := s.c.NewRequest(req)
	if err != nil {
		return nil, err
	}
	// Send request to Sirena
	response, err := s.c.Send(request)
	if err != nil {
		return nil, err
	}

	var availabilityResponse structs.AvailabilityResponse
	if err := xml.Unmarshal(response.Message, &availabilityResponse); err != nil {
		return nil, err
	}

	return &availabilityResponse, nil
}
