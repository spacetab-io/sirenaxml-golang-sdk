package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r *Repository) Refund(pnr, surname string, logAttributes map[string]string, pretend bool) (*sirena.RefundResponse, error) {
	sirenaRefundRequest := sirena.RefundRequest{
		Query: sirena.RefundRequestQuery{
			Refund: sirena.RefundRequestBody{
				Regnum:  pnr,
				Surname: surname,
				Action:  "query",
				Mode:    "refund",
				RequestParams: &sirena.RefundRequestParams{
					Pretend:          pretend,
					CheckForCash:     false,
					ShowPriceDetails: true,
					ShowPaydoc:       true,
				},
			},
		},
	}

	sirenaRefundRequestBytes, err := xml.Marshal(sirenaRefundRequest)
	if err != nil {
		return nil, err
	}

	sirenaRefundResponseXML, err := r.Request(sirenaRefundRequestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	// Decode Sirena response
	var sirenaRefundResponse sirena.RefundResponse
	if err := xml.Unmarshal(sirenaRefundResponseXML, &sirenaRefundResponse); err != nil {
		return nil, err
	}

	return &sirenaRefundResponse, nil
}
