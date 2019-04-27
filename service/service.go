package service

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/sdk"
)

type SirenaSDK interface {
	GetAvailability(req []byte) (*sdk.AvailabilityResponse, error)
}

type Service interface {
	Avalability(req *sdk.AvailabilityRequest) (*sdk.AvailabilityResponse, error)
}

type service struct {
	sdk SirenaSDK
}

func (s *service) Avalability(req *sdk.AvailabilityRequest) (*sdk.AvailabilityResponse, error) {
	availabiliteReqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	return s.sdk.GetAvailability(availabiliteReqXML)
}

func NewSKD(sdk SirenaSDK) Service {
	return &service{sdk: sdk}
}
