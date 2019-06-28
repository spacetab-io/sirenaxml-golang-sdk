package structs

import "encoding/xml"

type AvailabilityRequest struct {
	XMLName xml.Name                 `xml:"sirena"`
	Query   AvailabilityRequestQuery `xml:"query"`
}

type AvailabilityRequestQuery struct {
	Availability Availability `xml:"availability"`
}

type AvailabilityAnswerParams struct {
	ShowFlighttime bool `xml:"show_flighttime"`
}

type Availability struct {
	Departure    string                   `xml:"departure,omitempty"`
	Arrival      string                   `xml:"arrival,omitempty"`
	AnswerParams AvailabilityAnswerParams `xml:"answer_params,omitempty"`
	Flights      []AvailabilityFlight     `xml:"flight,omitempty"`
}

type AvailabilityResponse struct {
	XMLName xml.Name `xml:"sirena"`
	Answer  struct {
		Availability Availability `xml:"availability"`
	} `xml:"answer"`
}

type AvailabilityFlight struct {
	//XMLName        xml.Name `xml:"flight"`
	Num            int    `xml:"num"`
	Origin         string `xml:"origin"`
	OriginTerminal string `xml:"orig_term"`
	// .... @TODO расписать дальше!
}

type AvailabilityFlights struct {
}
