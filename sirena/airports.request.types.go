package sirena

import "encoding/xml"

// AirportsRequest is a <describe> request
type AirportsRequest struct {
	Query   AirportsRequestQuery `xml:"query"`
	XMLName xml.Name             `xml:"sirena"`
}

// AirportsRequestQuery is a <query> section in <describe> request
type AirportsRequestQuery struct {
	Airports Airports `xml:"describe"`
}

type Airports struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
