package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r *Repository) CheckAvailability(departure, arrival string, subclasses []string, logAttributes map[string]string) (*sirena.AvailabilityResponse, error) {

	requestParams := sirena.RequestParams{
		UseDag: false,
		UseIak: false,
	}

	availabilityRequest := sirena.AvailabilityRequestQuery{
		Departure:     departure,
		Arrival:       arrival,
		Subclass:      subclasses,
		RequestParams: requestParams,
	}

	requestBytes, err := xml.MarshalIndent(&availabilityRequest, "  ", "    ")
	if err != nil {
		return nil, err
	}

	responseBytes, err := r.Transport.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var availabilityResponse sirena.AvailabilityResponse
	if err = xml.Unmarshal(responseBytes, &availabilityResponse); err != nil {
		return nil, err
	}

	return &availabilityResponse, nil
}
