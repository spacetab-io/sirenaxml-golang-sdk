package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r *Repository) Fares(departure, arrival, passenger, company string, subclass []string, logAttributes map[string]string) (*sirena.FaresResponse, error) {
	query := sirena.FareRequestQuery{
		Fares: sirena.FaresQuery{
			Departure: departure,
			Arrival:   arrival,
			Company:   company,
			Subclass:  subclass,
			Passenger: passenger,
		},
	}
	request := sirena.FareRequest{
		Query: query,
	}

	requestBytes, err := xml.Marshal(request)
	if err != nil {
		// logs.Log.Error(err)
		// respond.With(w, r, http.StatusBadRequest, error.Error(err))
		return nil, err
	}

	faresResponseXML, err := r.Request(requestBytes, logAttributes)
	if err != nil {
		// logs.Log.Error(err)
		// respond.With(w, r, http.StatusBadRequest, error.Error(err))
		return nil, err
	}

	var faresResponse sirena.FaresResponse
	err = xml.Unmarshal(faresResponseXML, &faresResponse)
	if err != nil {
		// logs.Log.Error(err)
		// respond.With(w, r, http.StatusBadRequest, error.Error(err))
		return nil, err
	}

	return &faresResponse, nil
}
