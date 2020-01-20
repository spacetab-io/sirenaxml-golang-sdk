package repository

import (
	"encoding/xml"
	"strings"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r Repository) Order(logAttributes map[string]string, pnr string, surname string) (*sirena.OrderResponse, error) {

	sirenaOrderRequest := sirena.OrderRequest{
		Query: sirena.OrderRequestQuery{
			Order: sirena.Order{
				Regnum:  pnr,
				Surname: strings.ToUpper(surname),
				AnswerParams: sirena.OrderAnswerParams{
					Lang:            "en",
					AddRemoteRecloc: true,
				},
			},
		},
	}

	sirenaOrderRequestBytes, err := xml.MarshalIndent(&sirenaOrderRequest, "  ", "    ")
	if err != nil {
		return nil, err
	}

	sirenaOrderResponseXML, err := r.Request(sirenaOrderRequestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var sirenaOrderResponse sirena.OrderResponse
	if err := xml.Unmarshal(sirenaOrderResponseXML, &sirenaOrderResponse); err != nil {
		return nil, err
	}

	return &sirenaOrderResponse, nil
}
