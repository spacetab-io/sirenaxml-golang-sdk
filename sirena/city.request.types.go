package sirena

import "encoding/xml"

// CitiesRequest is a <describe> request
type CitiesRequest struct {
	Query   CitiesRequestQuery `xml:"query"`
	XMLName xml.Name           `xml:"sirena"`
}

// CitiesRequestQuery is a <query> section in <describe> request
type CitiesRequestQuery struct {
	Cities Cities `xml:"describe"`
}

type Cities struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
