package sirena

import "encoding/xml"

// CountriesRequest is a <describe> request
type CountriesRequest struct {
	Query   CountriesRequestQuery `xml:"query"`
	XMLName xml.Name              `xml:"sirena"`
}

// CountriesRequestQuery is a <query> section in <describe> request
type CountriesRequestQuery struct {
	Countries Countries `xml:"describe"`
}

type Countries struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
