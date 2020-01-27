package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r *Repository) ModifyPNRAdd(pnr, surname string, logAttributes map[string]string, version int, contacts []sirena.ModifyPNRContact, sirenaPassDocuments []sirena.ModifyPNRPassDocument) (*sirena.ModifyPNRResponse, error) {

	modifyPNRreq := sirena.ModifyPNRRequest{
		Query: sirena.ModifyPNRQuery{
			ModifyPNR: sirena.ModifyPNR{
				Regnum:  pnr,
				Surname: surname,
				Version: version,
				AddParams: sirena.ModifyPNRAddParams{
					Contact:      contacts,
					PassDocument: sirenaPassDocuments,
				},
			},
		},
	}

	modifyPNRrequest, err := xml.Marshal(modifyPNRreq)
	if err != nil {
		return nil, err
	}

	responseBytes, err := r.Transport.Request(modifyPNRrequest, logAttributes)
	if err != nil {
		return nil, err
	}

	var sirenaModifyPNRResponse sirena.ModifyPNRResponse
	err = xml.Unmarshal(responseBytes, &sirenaModifyPNRResponse)
	if err != nil {
		return nil, err
	}

	return &sirenaModifyPNRResponse, nil
}
