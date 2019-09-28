package structs

import "encoding/xml"

// RefundResponse is a Sirena response to <payment-ext-auth:refund> request
type RefundResponse struct {
	Answer  RefundAnswer `xml:"answer"`
	XMLName xml.Name     `xml:"sirena" json:"-"`
}

// RefundAnswer is an <answer> section in Sirena <payment-ext-auth:refund> response
type RefundAnswer struct {
	Pult   string               `xml:"pult,attr,omitempty"`
	Refund RefundAnswerResponse `xml:"payment-ext-auth"`
}

// RefundAnswerResponse is an <payment-ext-auth> section in Sirena <payment-ext-auth:refund> response
type RefundAnswerResponse struct {
	PNR    RefundAnswerPNR `xml:"pnr"`
	Regnum string          `xml:"regnum,attr,omitempty"`
	Action string          `xml:"action,attr,omitempty"`
	Mode   string          `xml:"mode,attr,omitempty"`
	Cost   struct {
		Currency string  `xml:"curr,attr"`
		Value    float64 `xml:",chardata"`
	} `xml:"cost"`
	Ok    *struct{} `xml:"ok,omitempty"`
	Error *Error    `xml:"error,omitempty"`
}

// RefundAnswerPNR is a <pnr> section in Sirena <payment-ext-auth:refund> response
type RefundAnswerPNR struct {
	Prices []RefundAnswerPNRPrice `xml:"prices>price"`
}

// RefundAnswerPNRPrice is a <price> subsection in Sirena <payment-ext-auth:refund> response
type RefundAnswerPNRPrice struct {
	SegmentID   int                      `xml:"segment-id,attr,omitempty"`
	PassengerID int                      `xml:"passenger-id,attr,omitempty"`
	Currency    string                   `xml:"currency,attr,omitempty"`
	Ticket      string                   `xml:"ticket,attr"`
	AcCode      string                   `xml:"accode,attr"`
	Fare        PayRefAnswerPNRPriceFare `xml:"fare"`
	Total       float64                  `xml:"total"`
	Taxes       RefundTaxes              `xml:"taxes"`
}

type RefundTaxes struct {
	Tax []Tax `xml:"tax"`
}

type Tax struct {
	Code  string `xml:"code"`
	Value Value  `xml:"value"`
}

type Value struct {
	Currency string `xml:"currency"`
	Value    string `xml:",chardata"`
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
