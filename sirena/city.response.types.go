package sirena

import "encoding/xml"

type CitiesResponse struct {
	Answer  CitiesAnswer `xml:"answer"`
	XMLName xml.Name     `xml:"sirena"`
}

type CitiesAnswer struct {
	Cities CitiesAnswerDetails `xml:"describe"`
}

type CitiesAnswerDetails struct {
	Data []CitiesAnswerData `xml:"data"`
}

type CitiesAnswerData struct {
	Code    []CitiesAnswerDataCode    `xml:"code"`
	Name    []CitiesAnswerDataName    `xml:"name"`
	Country []CitiesAnswerDataCountry `xml:"country"`
}

type CitiesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type CitiesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type CitiesAnswerDataCountry struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
