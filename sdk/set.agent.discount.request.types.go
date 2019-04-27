package sdk

import "encoding/xml"

// SetAgentDiscountRequest is a set_agent_discount request
type SetAgentDiscountRequest struct {
	Query   SetAgentDiscountQuery `xml:"query"`
	XMLName xml.Name              `xml:"sirena"`
}

// SetAgentDiscountQuery is a <query> section in set_agent_discount request
type SetAgentDiscountQuery struct {
	SetAgentDiscount *SetAgentDiscount `xml:"set_agent_discount"`
}

// SetAgentDiscount is a body of set_agent_discount request
type SetAgentDiscount struct {
	Regnum        *SetAgentDiscountRegnum        `xml:"regnum"`
	Unit          []*SetAgentDiscountUnit        `xml:"unit"`
	RequestParams *SetAgentDiscountRequestParams `xml:"request_params,omitempty"`
	AnswerParams  *BookingAnswerParams           `xml:"answer_params,omitempty"`
}

// SetAgentDiscountRegnum is a Regnum (PNR number and version) in set_agent_discount request
type SetAgentDiscountRegnum struct {
	Version int    `xml:"version,attr"`
	Value   string `xml:",chardata"`
}

// SetAgentDiscountUnit is a <unit> element of the set_agent_discount request
type SetAgentDiscountUnit struct {
	PassengerID          int                   `xml:"passenger-id,attr"`
	SegmentID            int                   `xml:"segment-id,attr"`
	FC                   string                `xml:"fc,attr"`
	SetAgentDiscountFare *SetAgentDiscountFare `xml:"fare"`
}

// SetAgentDiscountFare is a fare element in set_agent_discount request
type SetAgentDiscountFare struct {
	Discount string `xml:"discount,attr"`
	Brand    string `xml:"brand,attr,omitempty"`
	Value    string `xml:",chardata"`
}

// SetAgentDiscountRequestParams is a <request_params> section in set_agent_discount request
type SetAgentDiscountRequestParams struct {
	TickSer string `xml:"tick_ser"`
}
