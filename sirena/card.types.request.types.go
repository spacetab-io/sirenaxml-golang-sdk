package sirena

import "encoding/xml"

// CardTypesRequest is a Sirena <describe> request for all card types
type CardTypesRequest struct {
	Query   CardTypesRequestQuery `xml:"query"`
	XMLName xml.Name              `xml:"sirena"`
}

// CardTypesRequestQuery is a <query> section in  all card types request
type CardTypesRequestQuery struct {
	CardTypes CardTypes `xml:"describe"`
}

// CardTypes is a <describe> section in all card types request
type CardTypes struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
