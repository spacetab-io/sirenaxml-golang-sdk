package sirena

import "encoding/xml"

// AirlinesRequest is a <describe> request
type AirlinesRequest struct {
	Query   AirlinesRequestQuery `xml:"query"`
	XMLName xml.Name             `xml:"sirena"`
}

// AirlinesRequestQuery is a <query> section in <describe> request
type AirlinesRequestQuery struct {
	Airlines Airlines `xml:"describe"`
}

type Airlines struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
