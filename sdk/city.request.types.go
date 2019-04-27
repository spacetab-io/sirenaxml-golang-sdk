package sdk

import "encoding/xml"

// CitiesRequest is a Sirena <describe> request for all cities
type CitiesRequest struct {
	Query   CitiesRequestQuery `xml:"query"`
	XMLName xml.Name           `xml:"sirena"`
}

// CitiesRequestQuery is a <query> section in all cities request
type CitiesRequestQuery struct {
	Cities Cities `xml:"describe"`
}

// Cities is a <describe> section in all cities request
type Cities struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
