package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *Repository) ModifyPNRRemove(pnr, surname string, logAttributes map[string]string, keys []sirena.Ssr, version int) (*sirena.ModifyPNRResponse, error) {

	modifyPNRreq := sirena.ModifyPNRRequest{
		Query: sirena.ModifyPNRQuery{
			ModifyPNR: sirena.ModifyPNR{
				Regnum:  pnr,
				Surname: surname,
				Version: version,
				RemoveParams: sirena.ModifyPNRRemoveParams{
					Ssr: keys,
				},
			},
		},
	}

	modifyPNRrequest, err := xml.Marshal(modifyPNRreq)
	if err != nil {
		return nil, err
	}

	responseBytes, err := s.Request(modifyPNRrequest, logAttributes)
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
