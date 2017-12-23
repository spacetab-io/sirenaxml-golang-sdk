package sirena

import "encoding/xml"

type AirportsResponse struct {
	Answer  AirportsAnswer `xml:"answer"`
	XMLName xml.Name       `xml:"sirena" json:"-"`
}

type AirportsAnswer struct {
	Airports AirportsAnswerDetails `xml:"describe"`
}

type AirportsAnswerDetails struct {
	Data []AirportsAnswerData `xml:"data"`
}

type AirportsAnswerData struct {
	Code []AirportsAnswerDataCode `xml:"code"`
	Name []AirportsAnswerDataName `xml:"name"`
	City []AirportsAnswerDataCity `xml:"city"`
}

type AirportsAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type AirportsAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type AirportsAnswerDataCity struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
