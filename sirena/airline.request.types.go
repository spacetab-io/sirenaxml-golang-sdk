package sirena

import "encoding/xml"

// AirlinesRequest is a Sirena <describe> request for all airlines
type AirlinesRequest struct {
	Query   AirlinesRequestQuery `xml:"query"`
	XMLName xml.Name             `xml:"sirena"`
}

// AirlinesRequestQuery is a <query> section in all airlines request
type AirlinesRequestQuery struct {
	Airlines Airlines `xml:"describe"`
}

// Airlines is a <describe> section in all airlines request
type Airlines struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
