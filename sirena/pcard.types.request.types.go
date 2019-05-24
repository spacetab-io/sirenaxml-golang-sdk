package sirena

import "encoding/xml"

// PCardTypesRequest is a Sirena <describe> request for all pcard_types
type PCardTypesRequest struct {
	Query   PCardTypesRequestQuery `xml:"query"`
	XMLName xml.Name               `xml:"sirena"`
}

// PCardTypesRequestQuery is a <query> section in  all pcard_types request
type PCardTypesRequestQuery struct {
	PCardTypes PCardTypes `xml:"describe"`
}

// PCardTypes is a <describe> section in all pcard_types request
type PCardTypes struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
