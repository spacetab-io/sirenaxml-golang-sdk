package repository

import (
	"encoding/xml"
	"fmt"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s *Repository) RefundConfirm(pnr string, logAttributes map[string]string, totalCost float64, surname string) (*sirena.RefundResponse, error) {
	sirenaRefundRequest := sirena.RefundRequest{
		Query: sirena.RefundRequestQuery{
			Refund: sirena.RefundRequestBody{
				Regnum:  pnr,
				Surname: surname,
				Action:  "confirm",
				Mode:    "refund",
				Cost: &sirena.Cost{
					Curr:  "RUB",
					Value: fmt.Sprintf("%.2f", totalCost),
				},
			},
		},
	}

	sirenaRefundRequestBytes, err := xml.Marshal(sirenaRefundRequest)
	if err != nil {
		return nil, err
	}

	sirenaRefundResponseXML, err := s.Request(sirenaRefundRequestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	// Decode Sirena response
	var sirenaRefundResponse sirena.RefundResponse
	err = xml.Unmarshal(sirenaRefundResponseXML, &sirenaRefundResponse)
	if 	err != nil {
		return nil, err
	}

	return &sirenaRefundResponse, nil
}
