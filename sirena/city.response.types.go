package sirena

import "encoding/xml"

// CitiesResponse is a response to all cities request
type CitiesResponse struct {
	Answer  CitiesAnswer `xml:"answer"`
	XMLName xml.Name     `xml:"sirena" json:"-"`
}

// CitiesAnswer is an <answer> section in all cities response
type CitiesAnswer struct {
	Cities CitiesAnswerDetails `xml:"describe"`
}

// CitiesAnswerDetails is a <describe> section in all cities response
type CitiesAnswerDetails struct {
	Data []CitiesAnswerData `xml:"data"`
}

// CitiesAnswerData is a <data> section in all cities response
type CitiesAnswerData struct {
	Code    []CitiesAnswerDataCode    `xml:"code"`
	Name    []CitiesAnswerDataName    `xml:"name"`
	Country []CitiesAnswerDataCountry `xml:"country"`
}

// CitiesAnswerDataCode represents <code> entry in <data> section
type CitiesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// CitiesAnswerDataName represents <name> entry in <data> section
type CitiesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// CitiesAnswerDataCountry represents <country> entry in <data> section
type CitiesAnswerDataCountry struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
