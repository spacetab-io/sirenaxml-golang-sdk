package structs

import "encoding/xml"

// AirportsRequest is a Sirena <describe> request for all airports
type AirportsRequest struct {
	Query   AirportsRequestQuery `xml:"query"`
	XMLName xml.Name             `xml:"sirena"`
}

// AirportsRequestQuery is a <query> section in all airports request
type AirportsRequestQuery struct {
	Airports Airports `xml:"describe"`
}

// Airports is a <describe> section in all airports request
type Airports struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
