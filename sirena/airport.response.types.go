package sirena

import "encoding/xml"

// AirportsResponse is a response to all airports request
type AirportsResponse struct {
	Answer  AirportsAnswer `xml:"answer"`
	XMLName xml.Name       `xml:"sirena" json:"-"`
}

// AirportsAnswer is an <answer> section in all airports response
type AirportsAnswer struct {
	Airports AirportsAnswerDetails `xml:"describe"`
}

// AirportsAnswerDetails is a <describe> section in all airports response
type AirportsAnswerDetails struct {
	Data []AirportsAnswerData `xml:"data"`
}

// AirlinesAnswerData is a <data> section in all airports response
type AirportsAnswerData struct {
	Code []AirportsAnswerDataCode `xml:"code"`
	Name []AirportsAnswerDataName `xml:"name"`
	City []AirportsAnswerDataCity `xml:"city"`
}

// AirportsAnswerDataCode represents <code> entry in <data> section
type AirportsAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// AirportsAnswerDataName represents <name> entry in <data> section
type AirportsAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// AirportsAnswerDataCity represents <city> entry in <data> section
type AirportsAnswerDataCity struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
