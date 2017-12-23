package sirena

import "encoding/xml"

type CountriesResponse struct {
	Answer  CountriesAnswer `xml:"answer"`
	XMLName xml.Name        `xml:"sirena" json:"-"`
}

type CountriesAnswer struct {
	Countries CountriesAnswerDetails `xml:"describe"`
}

type CountriesAnswerDetails struct {
	Data []CountriesAnswerData `xml:"data"`
}

type CountriesAnswerData struct {
	Code []CountriesAnswerDataCode `xml:"code"`
	Name []CountriesAnswerDataName `xml:"name"`
}

type CountriesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type CountriesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
