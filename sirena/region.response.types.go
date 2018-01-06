package sirena

import "encoding/xml"

// RegionsResponse is a response to all regions request
type RegionsResponse struct {
	Answer  RegionsAnswer `xml:"answer"`
	XMLName xml.Name      `xml:"sirena" json:"-"`
}

// RegionsAnswer is an <answer> section in all regions response
type RegionsAnswer struct {
	Regions RegionsAnswerDetails `xml:"describe"`
}

// RegionsAnswerDetails is a <describe> section in all regions response
type RegionsAnswerDetails struct {
	Data []RegionsAnswerData `xml:"data"`
}

// RegionsAnswerData is a <data> section in all regions response
type RegionsAnswerData struct {
	Code    []RegionsAnswerDataCode    `xml:"code"`
	Name    []RegionsAnswerDataName    `xml:"name"`
	Country []RegionsAnswerDataCountry `xml:"country"`
}

// RegionsAnswerDataCode represents <code> entry in <data> section
type RegionsAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// RegionsAnswerDataName represents <name> entry in <data> section
type RegionsAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// RegionsAnswerDataCountry represents <country> entry in <data> section
type RegionsAnswerDataCountry struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
