package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r Repository) SetAgentDiscount(logAttributes map[string]string, version int, pnr string, units []*sirena.SetAgentDiscountUnit) (*sirena.SetAgentDiscountResponse, error) {

	sirenaSetAgentDiscountRequest := &sirena.SetAgentDiscountRequest{
		Query: sirena.SetAgentDiscountQuery{
			SetAgentDiscount: &sirena.SetAgentDiscount{
				Regnum: &sirena.SetAgentDiscountRegnum{
					Version: version,
					Value:   pnr,
				},
				Unit: units,
				RequestParams: &sirena.SetAgentDiscountRequestParams{
					TickSer: "ETM",
				},
				AnswerParams: &sirena.BookingAnswerParams{
					Lang: "en",
				},
			},
		},
	}

	setAgentDiscountRequestBytes, err := xml.MarshalIndent(&sirenaSetAgentDiscountRequest, "  ", "    ")
	if err != nil {
		return nil, err
	}

	setAgentDiscountResponseXML, err := r.Transport.Request(setAgentDiscountRequestBytes, logAttributes)
	if err != nil {
		return nil, err

	}

	var sirenaSetAgentDiscountResponse sirena.SetAgentDiscountResponse
	if err := xml.Unmarshal(setAgentDiscountResponseXML, &sirenaSetAgentDiscountResponse); err != nil {
		return nil, err
	}

	return &sirenaSetAgentDiscountResponse, nil
}
