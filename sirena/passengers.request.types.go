package sirena

import "encoding/xml"

// PassengersRequest is a Sirena <describe> request for passenger(s)
type PassengersRequest struct {
	Query   PassengersRequestQuery `xml:"query"`
	XMLName xml.Name               `xml:"sirena"`
}

// PassengersRequestQuery is a <query> section in passenger(s) request
type PassengersRequestQuery struct {
	Passengers Passengers `xml:"describe"`
}

// Passengers is a <describe> section in passenger(s) request
type Passengers struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
