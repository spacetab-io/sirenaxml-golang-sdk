package service

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *service) Availability(req *structs.AvailabilityRequest) (*structs.Availability, error) {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.sdk.GetAvailability(reqXML)
	if err != nil {
		return nil, err
	}

	return &resp.Answer.Availability, nil
}
