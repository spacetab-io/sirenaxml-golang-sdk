package sirena

import "encoding/xml"

// CountriesResponse is a response to all countries request
type CountriesResponse struct {
	Answer  CountriesAnswer `xml:"answer"`
	XMLName xml.Name        `xml:"sirena" json:"-"`
}

// CountriesAnswer is an <answer> section in all countries response
type CountriesAnswer struct {
	Countries CountriesAnswerDetails `xml:"describe"`
}

// CountriesAnswerDetails is a <describe> section in all countries response
type CountriesAnswerDetails struct {
	Data []CountriesAnswerData `xml:"data"`
}

// CountriesAnswerData is a <data> section in all countries response
type CountriesAnswerData struct {
	Code []CountriesAnswerDataCode `xml:"code"`
	Name []CountriesAnswerDataName `xml:"name"`
}

// CountriesAnswerDataCode represents <code> entry in <data> section
type CountriesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// CountriesAnswerDataName represents <name> entry in <data> section
type CountriesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
