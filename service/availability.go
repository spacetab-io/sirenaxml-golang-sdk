package service

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *service) Avalability(req *structs.AvailabilityRequest) (*structs.AvailabilityResponse, error) {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	return s.sdk.GetAvailability(reqXML)
}
