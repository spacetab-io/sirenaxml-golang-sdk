package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s Repository) GetItinReceipts(logAttributes map[string]string, pnr string, surname string) (*sirena.GetItinReceiptsDataResponse, error) {

	sirenaGetItinReceiptsDataRequest := sirena.GetItinReceiptsDataRequest{
		Query: sirena.GetItinReceiptsDataRequestQuery{
			GetItinReceiptsData: sirena.GetItinReceiptsData{
				Regnum:  pnr,
				Surname: surname,
			},
		},
	}

	requestBytes, err := xml.Marshal(sirenaGetItinReceiptsDataRequest)
	if err != nil {
		// logs.Log.Error(nil, err)
		return nil, err
	}

	responseBytes, err := s.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var sirenaGetItinReceiptsDataResponse sirena.GetItinReceiptsDataResponse
	if err := xml.Unmarshal(responseBytes, &sirenaGetItinReceiptsDataResponse); err != nil {
		return nil, err
	}

	return &sirenaGetItinReceiptsDataResponse, nil
}
