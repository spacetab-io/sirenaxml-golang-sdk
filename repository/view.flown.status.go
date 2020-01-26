package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r *Repository) ViewFlownStatus(surname, pnr string, logAttributes map[string]string) (*sirena.ViewFlownStatusResponse, error) {

	ViewFlownStatusRequest := sirena.ViewFlownStatusRequest{
		Query: sirena.ViewFlownStatusQuery{
			ViewFlownStatus: sirena.ViewFlownStatus{
				Regnum:  pnr,
				Surname: surname,
			},
		},
	}

	viewFlownStatusRequestBytes, err := xml.Marshal(ViewFlownStatusRequest)
	if err != nil {
		return nil, err
	}

	viewFlownStatusResponseXML, err := r.Transport.Request(viewFlownStatusRequestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var sirenaViewFlownStatusResponse sirena.ViewFlownStatusResponse
	if err := xml.Unmarshal(viewFlownStatusResponseXML, &sirenaViewFlownStatusResponse); err != nil {
		return nil, err
	}

	return &sirenaViewFlownStatusResponse, nil
}
