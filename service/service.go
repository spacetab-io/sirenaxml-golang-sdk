package service

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

type SirenaSDK interface {
	GetAvailability(req []byte) (*structs.AvailabilityResponse, error)
}

type Service interface {
	Avalability(req *structs.AvailabilityRequest) (*structs.AvailabilityResponse, error)
}

type service struct {
	sdk SirenaSDK
}

func (s *service) Avalability(req *structs.AvailabilityRequest) (*structs.AvailabilityResponse, error) {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	return s.sdk.GetAvailability(reqXML)
}

func NewSKD(sdk SirenaSDK) Service {
	return &service{sdk: sdk}
}
