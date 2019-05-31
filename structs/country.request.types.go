package structs

import "encoding/xml"

// CountriesRequest is a Sirena <describe> request for all countries
type CountriesRequest struct {
	Query   CountriesRequestQuery `xml:"query"`
	XMLName xml.Name              `xml:"sirena"`
}

// CountriesRequestQuery is a <query> section in <describe> request
type CountriesRequestQuery struct {
	Countries Countries `xml:"describe"`
}

// Countries is a <describe> section in all countries request
type Countries struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
