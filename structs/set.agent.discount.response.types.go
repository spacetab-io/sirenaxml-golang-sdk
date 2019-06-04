package structs

import (
	"encoding/xml"
)

// SetAgentDiscountResponse is a Sirena response to set_agent_discount request
type SetAgentDiscountResponse struct {
	Answer  SetAgentDiscountAnswer `xml:"answer"`
	XMLName xml.Name               `xml:"sirena" json:"-"`
}

// SetAgentDiscountAnswer is an <answer> section in Sirena set_agent_discount response
type SetAgentDiscountAnswer struct {
	SetAgentDiscount *SetAgentDiscountAnswerBody `xml:"set_agent_discount"`
}

// SetAgentDiscountAnswerBody is an <set_agent_discount> section in Sirena set_agent_discount response
type SetAgentDiscountAnswerBody struct {
	Regnum string                     `xml:"regnum,attr"`
	PNR    *SetAgentDiscountAnswerPNR `xml:"pnr"`
	Error  *Error             `xml:"error,omitempty"`
}

// SetAgentDiscountAnswerPNR is a <pnr> section in Sirena set_agent_discount response
type SetAgentDiscountAnswerPNR struct {
	Prices *BookingAnswerPNRPrices `xml:"prices"`
}
