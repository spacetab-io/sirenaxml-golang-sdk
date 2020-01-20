package repository

import (
	"encoding/xml"
	"fmt"
	"strings"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r Repository) Payment(
	logAttributes map[string]string,
	surname string,
	pnr string,
	action string,
	paydoc sirena.Paydoc,
	currency string,
	cost float64,
	tickSer string,
	paymentTimeout int,
	returnReceipt bool,
) (
	*sirena.PaymentResponse,
	error,
) {
	sirenaPaymentRequest := sirena.PaymentRequest{
		Query: sirena.PaymentRequestQuery{
			Payment: sirena.PaymentRequestBody{
				Regnum:  pnr,
				Surname: strings.ToUpper(surname),
				Action:  action,
				Paydoc: sirena.Paydoc{
					Formpay:  paydoc.Formpay,
					Type:     paydoc.Type,
					Num:      paydoc.Num,
					ExpDate:  paydoc.ExpDate,
					Holder:   paydoc.Holder,
					AuthCode: paydoc.AuthCode,
				},
				Cost: &sirena.Cost{
					Curr:  currency,
					Value: fmt.Sprintf("%.2f", cost),
				},
				RequestParams: sirena.PaymentRequestParams{
					TickSer:        tickSer,
					PaymentTimeout: paymentTimeout,
					ReturnReceipt:  returnReceipt,
				},
			},
		},
	}

	requestBytes, err := xml.MarshalIndent(&sirenaPaymentRequest, "  ", "    ")
	if err != nil {
		return nil, err
	}

	responseBytes, err := r.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var paymentResponse sirena.PaymentResponse
	if err = xml.Unmarshal(responseBytes, &paymentResponse); err != nil {
		return nil, err
	}

	return &paymentResponse, nil
}
