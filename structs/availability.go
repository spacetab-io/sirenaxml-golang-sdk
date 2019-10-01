package structs

import "encoding/xml"

type AvailabilityRequest struct {
	XMLName xml.Name                 `xml:"sirena"`
	Query   AvailabilityRequestQuery `xml:"query"`
}

type AvailabilityRequestQuery struct {
	Departure     string        `xml:"departure,omitempty"`
	Arrival       string        `xml:"arrival,omitempty"`
	Subclass      []string      `xml:"subclass"`
	RequestParams RequestParams `xml:"request_params"`
}

type RequestParams struct {
	UseDag               bool `xml:"use_dag"`
	UseIak               bool `xml:"use_iak"`
	CheckTchRestrictions bool `xml:"check_tch_restrictions"`
}

type AvailabilityAnswerParams struct {
	ShowFlightTime bool `xml:"show_flighttime"`
}

type Availability struct {
	Departure    string                   `xml:"departure,omitempty"`
	Arrival      string                   `xml:"arrival,omitempty"`
	AnswerParams AvailabilityAnswerParams `xml:"answer_params,omitempty"`
	Flight       AvailabilityFlight       `xml:"flight,omitempty"`
	Flights      AvailabilityFlights      `xml:"flights,omitempty"`
}

type AvailabilityResponse struct {
	XMLName xml.Name `xml:"sirena"`
	Answer  Answer   `xml:"answer"`
}

type Answer struct {
	Availability Availability `xml:"availability"`
}

type AvailabilityFlights struct {
	Flight []AvailabilityFlight `xml:"flight,omitempty"`
}

type AvailabilityFlight struct {
	Company        string   `xml:"company"`
	Destination    string   `xml:"destination"`
	Subclass       Subclass `xml:"subclass"`
	Num            int      `xml:"num"`
	Origin         string   `xml:"origin"`
	OriginTerminal string   `xml:"orig_term"`
}

type Subclass struct {
	Value string `xml:",chardata"`
	Count string `xml:"count,attr"`
}
