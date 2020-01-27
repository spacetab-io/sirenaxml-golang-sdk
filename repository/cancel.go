package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r *Repository) Cancel(pnr string, surname string, logAttributes map[string]string) (*sirena.OrderCancelResponse, error) {
	Query := sirena.CancelRequestQuery{
		Cancel: sirena.CancelRequestBody{
			Regnum:  pnr,
			Surname: surname,
		},
	}

	sirenaCancelRequest := sirena.CancelRequest{
		Query: Query,
	}

	requestBytes, err := xml.Marshal(sirenaCancelRequest)
	if err != nil {
		// logs.Log.Error(err)
		// respond.With(w, r, http.StatusBadRequest, error.Error(err))
		return nil, err
	}

	cancelResponseXML, err := r.Transport.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var orderCancelResponse sirena.OrderCancelResponse
	if err := xml.Unmarshal(cancelResponseXML, &orderCancelResponse); err != nil {
		return nil, err
	}

	return &orderCancelResponse, nil
}
