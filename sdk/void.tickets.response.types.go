package sdk

import "encoding/xml"

// VoidTicketsResponse is a Sirena response to <void_tickets> request
type VoidTicketsResponse struct {
	Answer  VoidTicketsAnswer `xml:"answer"`
	XMLName xml.Name          `xml:"sirena" json:"-"`
}

// VoidTicketsAnswer is an <answer> section in Sirena <void_tickets> response
type VoidTicketsAnswer struct {
	Pult        string                 `xml:"pult,attr,omitempty"`
	VoidTickets VoidTicketsAnswerQuery `xml:"void_tickets"`
}

// VoidTicketsAnswerQuery is an <order> section in Sirena order response
type VoidTicketsAnswerQuery struct {
	TicketsReturned bool           `xml:"tickets_returned"`
	Error           *ErrorResponse `xml:"error"`
}
