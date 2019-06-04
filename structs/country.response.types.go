package structs

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
	Data  []CountriesAnswerData `xml:"data"`
	Error *Error                `xml:"error,omitempty"`
}

// CountriesAnswerData is a <data> section in all countries response
type CountriesAnswerData struct {
	Code []CountriesAnswerDataCode `xml:"code" json:"code"`
	Name []CountriesAnswerDataName `xml:"name" json:"name"`
}

// CountriesAnswerDataCode represents <code> entry in <data> section
type CountriesAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// CountriesAnswerDataName represents <name> entry in <data> section
type CountriesAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
