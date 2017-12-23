package sirena

import "encoding/xml"

type RegionsResponse struct {
	Answer  RegionsAnswer `xml:"answer"`
	XMLName xml.Name      `xml:"sirena" json:"-"`
}

type RegionsAnswer struct {
	Regions RegionsAnswerDetails `xml:"describe"`
}

type RegionsAnswerDetails struct {
	Data []RegionsAnswerData `xml:"data"`
}

type RegionsAnswerData struct {
	Code    []RegionsAnswerDataCode    `xml:"code"`
	Name    []RegionsAnswerDataName    `xml:"name"`
	Country []RegionsAnswerDataCountry `xml:"country"`
}

type RegionsAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type RegionsAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type RegionsAnswerDataCountry struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
