package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s Repository) FareRemark(logAttributes map[string]string, fareRemarkRequest *sirena.FareRemarkRequest) (*sirena.FareRemarkResponse, error) {

	requestBytes, err := xml.MarshalIndent(&fareRemarkRequest, "  ", "    ")
	if err != nil {
		return nil, err
	}

	sirenaFareRemarkResponseXML, err := s.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var sirenaFareRemarkResponse sirena.FareRemarkResponse
	if err = xml.Unmarshal(sirenaFareRemarkResponseXML, &sirenaFareRemarkResponse); err != nil {
		return nil, err
	}

	return &sirenaFareRemarkResponse, nil
}
