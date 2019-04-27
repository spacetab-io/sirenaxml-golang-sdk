package sdk

import "encoding/xml"

// PaymentResponse is a Sirena response to <payment-ext-auth> request
type PaymentResponse struct {
	Answer  PaymentAnswer `xml:"answer"`
	XMLName xml.Name      `xml:"sirena" json:"-"`
}

// PaymentAnswer is an <answer> section in Sirena <payment-ext-auth> response
type PaymentAnswer struct {
	Pult    string             `xml:"pult,attr,omitempty"`
	Payment PaymentAnswerQuery `xml:"payment-ext-auth"`
}

// PaymentAnswerQuery is an <order> section in Sirena order response
type PaymentAnswerQuery struct {
	Cost struct {
		Currency string  `xml:"curr,attr"`
		Value    float64 `xml:",chardata"`
	} `xml:"cost"`
	Timeout          int                   `xml:"timeout"`
	Regnum           string                `xml:"regnum,omitempty"`
	NSeats           int                   `xml:"nseats,omitempty"`
	Agn              string                `xml:"agn,omitempty"`
	PPR              string                `xml:"ppr,omitempty"`
	Tickets          []PaymentAnswerTicket `xml:"tickinfo"`
	VoidTimeLimitUTC string                `xml:"void_timelimit_utc,omitempty"` // TimeDate format
	Error            *ErrorResponse        `xml:"error"`
}

// PaymentAnswerTicket holds ticket details in payment confirm response
type PaymentAnswerTicket struct {
	TickNum string `xml:"ticknum,attr"`
	PassID  string `xml:"pass_id,attr"`
	SegID   string `xml:"seg_id,attr"`
	ACCode  string `xml:"accode,attr"`
}
