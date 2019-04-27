package sdk

import (
	"encoding/xml"
)

// PricingMonobrandRequest is a <pricing_mono_brand> request
type PricingMonobrandRequest struct {
	Query   PricingMonobrandRequestQuery `xml:"query"`
	XMLName xml.Name                     `xml:"sirena"`
}

// PricingMonobrandRequestQuery is a <query> section in <pricing_mono_brand> request
type PricingMonobrandRequestQuery struct {
	PricingMonobrand PricingMonobrand `xml:"pricing_mono_brand"`
}

// PricingMonobrand is a <pricing_mono_brand> entry in <pricing_mono_brand> response
type PricingMonobrand struct {
	Segments      []PricingRequestSegment   `xml:"segment"`
	Passenger     []PricingRequestPassenger `xml:"passenger"`
	RequestParams PricingRequestParams      `xml:"request_params"`
	AnswerParams  PricingAnswerParams       `xml:"answer_params"`
}
