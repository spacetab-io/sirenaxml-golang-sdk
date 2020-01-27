package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r Repository) AddRemark(logAttributes map[string]string, pnr, surname, remarkType, remark string) (*sirena.AddRemarkResponse, error) {

	sirenaAddRemarkRequest := sirena.AddRemarkRequest{
		Query: sirena.AddRemarkRequestQuery{
			AddRemark: sirena.AddRemark{
				Regnum:  pnr,
				Surname: surname,
				Type:    remarkType,
				Remark:  remark,
			},
		},
	}

	addRemarkRequestBytes, err := xml.MarshalIndent(&sirenaAddRemarkRequest, "  ", "    ")
	if err != nil {
		return nil, err
	}

	sirenaAddRemarkResponseXML, err := r.Transport.Request(addRemarkRequestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var sirenaAddRemarkResponse sirena.AddRemarkResponse
	if err := xml.Unmarshal(sirenaAddRemarkResponseXML, &sirenaAddRemarkResponse); err != nil {
		return nil, err
	}

	return &sirenaAddRemarkResponse, nil
}
