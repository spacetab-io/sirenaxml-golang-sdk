package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
	// error2 "gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/sirena-agent-go/errorcode"
	// "gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/sirena-agent-go/http/structs"
	// "gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/sirena-agent-go/logs"
	// "gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/sirena-agent-go/pkg/models"
	// "gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/sirena-agent-go/publisher"
)

func (r *Repository) AddFFInfo(pnr, surname string, sirenaRequestPassengers []sirena.AddFFInfoRequestPassenger, logAttributes map[string]string) (*sirena.AddFFInfoResponse, error) {

	sirenaAddFFInfoRequest := sirena.AddFFInfoRequest{
		Query: sirena.AddFFInfoRequestQuery{
			AddFFInfo: sirena.AddFFInfo{
				Regnum:    pnr,
				Surname:   surname,
				Passenger: sirenaRequestPassengers,
			},
		},
	}

	// Encode Sirena add_remark request into XML
	AddFFInfoRequest, err := xml.MarshalIndent(&sirenaAddFFInfoRequest, "  ", "    ")
	if err != nil {
		return nil, err
	}

	sirenaAddFFInfoResponseXML, err := r.Request(AddFFInfoRequest, logAttributes)
	if err != nil {
		return nil, err
	}

	// Decode Sirena response
	var sirenaAddRemarkResponse sirena.AddFFInfoResponse
	err = xml.Unmarshal(sirenaAddFFInfoResponseXML, &sirenaAddRemarkResponse)
	if err != nil {
		return nil, err
	}

	return &sirenaAddRemarkResponse, nil
}
