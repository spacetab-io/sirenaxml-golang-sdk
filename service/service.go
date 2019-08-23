package service

import (
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/message"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

type Storage interface {
	SendRawRequest(req []byte) (resp []byte, err error)
	GetAvailability(req []byte) (*structs.AvailabilityResponse, error)
	GetCurrentKeyInfo(req []byte) (*structs.KeyInfoResponse, error)
	GetKeyData() (*message.KeyData, error)
}

type Service interface {
	RawRequest(req []byte) ([]byte, error)
	Availability(req *structs.AvailabilityRequest) (*structs.Availability, *structs.Error, error)
	KeyInfo() (*structs.KeyManager, *structs.Error, error)
}

type service struct {
	sdk Storage
}

func New(sdk Storage) Service {
	return &service{sdk: sdk}
}

func (s *service) RawRequest(req []byte) ([]byte, error) {
	data, err := s.sdk.SendRawRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}
