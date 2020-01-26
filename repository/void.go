package repository

import (
	"encoding/xml"

	"github.com/pkg/errors"
	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r Repository) Void(logAttributes map[string]string, pnr string, surname string) (*sirena.VoidTicketsResponse, error) {

	sirenaVoidTicketsRequest := sirena.VoidTicketsRequest{
		Query: sirena.VoidTicketsRequestQuery{
			VoidTickets: sirena.VoidTickets{
				Regnum:  pnr,
				Surname: surname,
				RequestParams: sirena.VoidTicketsRequestParams{
					ReturnSeats: true,
				},
			},
		},
	}

	requestBytes, err := xml.MarshalIndent(&sirenaVoidTicketsRequest, "  ", "    ")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	responseBytes, err := r.Transport.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var sirenaVoidResponse sirena.VoidTicketsResponse

	err = xml.Unmarshal(responseBytes, &sirenaVoidResponse)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sirenaVoidResponse, err
}
