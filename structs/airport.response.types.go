package structs

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
	Data  []AirportsAnswerData `xml:"data"`
	Error *Error               `xml:"error"`
}

// AirlinesAnswerData is a <data> section in all airports response
type AirportsAnswerData struct {
	Code []AirportsAnswerDataCode `xml:"code" json:"code"`
	Name []AirportsAnswerDataName `xml:"name" json:"name"`
	City []AirportsAnswerDataCity `xml:"city" json:"city"`
}

// AirportsAnswerDataCode represents <code> entry in <data> section
type AirportsAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// AirportsAnswerDataName represents <name> entry in <data> section
type AirportsAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// AirportsAnswerDataCity represents <city> entry in <data> section
type AirportsAnswerDataCity struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
