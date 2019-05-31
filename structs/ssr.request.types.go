package structs

import "encoding/xml"

// SSRRequest is a Sirena <describe> request for all special requests
type SSRRequest struct {
	Query   SSRRequestQuery `xml:"query"`
	XMLName xml.Name        `xml:"sirena"`
}

// SSRRequestQuery is a <query> section in  all special requests request
type SSRRequestQuery struct {
	SSR SSR `xml:"describe"`
}

// SSR is a <describe> section in all special requests request
type SSR struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
