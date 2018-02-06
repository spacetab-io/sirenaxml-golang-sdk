package sirena

import "encoding/xml"

// PaymentRefundResponse is a Sirena response to <payment-ext-auth:refund> request
type PaymentRefundResponse struct {
	Answer  PaymentRefundAnswer `xml:"answer"`
	XMLName xml.Name            `xml:"sirena" json:"-"`
}

// PaymentRefundAnswer is an <answer> section in Sirena <payment-ext-auth:refund> response
type PaymentRefundAnswer struct {
	Pult          string                      `xml:"pult,attr,omitempty"`
	PaymentRefund PaymentRefundAnswerResponse `xml:"payment-ext-auth"`
}

// PaymentRefundAnswerResponse is an <payment-ext-auth> section in Sirena <payment-ext-auth:refund> response
type PaymentRefundAnswerResponse struct {
	PNR    PaymentRefundAnswerPNR `xml:"pnr"`
	Regnum string                 `xml:"regnum,attr,omitempty"`
	Action string                 `xml:"action,attr,omitempty"`
	Mode   string                 `xml:"mode,attr,omitempty"`
	Cost   struct {
		Currency string  `xml:"curr,attr"`
		Value    float64 `xml:",chardata"`
	} `xml:"cost"`
	Error *Error `xml:"error"`
}

// PaymentRefundAnswerPNR is a <pnr> section in Sirena <payment-ext-auth:refund> response
type PaymentRefundAnswerPNR struct {
	Prices []PaymentRefundAnswerPNRPrice `xml:"prices>price"`
}

// PaymentRefundAnswerPNRPrice is a <price> subsection in Sirena <payment-ext-auth:refund> response
type PaymentRefundAnswerPNRPrice struct {
	SegmentID   int                      `xml:"segment-id,attr,omitempty"`
	PassengerID int                      `xml:"passenger-id,attr,omitempty"`
	Fare        PayRefAnswerPNRPriceFare `xml:"fare"`
}

// PayRefAnswerPNRPriceFare is a <fare> subsection in Sirena <payment-ext-auth:refund> response
type PayRefAnswerPNRPriceFare struct {
	FareExpdate string                        `xml:"fare_expdate,attr,omitempty"` // DateTime format
	Code        PayRefAnswerPNRPriceFareCode  `xml:"code"`
	Value       PayRefAnswerPNRPriceFareValue `xml:"value"`
}

// PayRefAnswerPNRPriceFareCode is a <code> element of a <fare> subsection in Sirena <payment-ext-auth:refund> response
type PayRefAnswerPNRPriceFareCode struct {
	BaseCode string `xml:"base_code,attr,omitempty"`
	Value    string `xml:",chardata"`
}

// PayRefAnswerPNRPriceFareValue is a <value> element of a <fare> subsection in Sirena <payment-ext-auth:refund> response
type PayRefAnswerPNRPriceFareValue struct {
	Currency string  `xml:"currency,attr,omitempty"`
	Value    float64 `xml:",chardata"`
}
