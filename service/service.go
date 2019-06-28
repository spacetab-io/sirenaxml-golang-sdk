package service

import (
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/client"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

type Storage interface {
	SendRawRequest(req []byte) ([]byte, error)
	GetAvailability(req []byte) (*structs.AvailabilityResponse, error)
	GetCurrentKeyInfo(req []byte) (*structs.KeyInfoResponse, error)
	GetKeyData() (*client.KeyData, error)
}

type Service interface {
	RawRequest(req []byte) ([]byte, error)
	Availability(req *structs.AvailabilityRequest) (*structs.Availability, error)
	KeyInfo() (*structs.KeyManager, error)
}

type service struct {
	sdk Storage
}

func NewSKD(sdk Storage) Service {
	return &service{sdk: sdk}
}

func (s *service) RawRequest(req []byte) ([]byte, error) {
	return s.sdk.SendRawRequest(req)
}
