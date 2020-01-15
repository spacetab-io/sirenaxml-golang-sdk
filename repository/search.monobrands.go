package repository

import (
	"encoding/xml"

	"github.com/pkg/errors"
	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

// @TODO возвращать структуру, а не байты. Потом делать анмаршал будет коряво
func (s Repository) SearchMonobrands(logAttributes map[string]string, request sirena.PricingMonobrandRequest) ([]byte, error) {

	requestBytes, err := xml.MarshalIndent(request, "  ", "    ")
	if err != nil {
		return nil, errors.Wrap(err, "pricingMonobrandRequest marshal error")
	}

	monobrandsResp, err := s.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	return monobrandsResp, nil
}
