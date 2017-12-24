package sirena

import "encoding/xml"

// CardTypesResponse is a response to all card types request
type CardTypesResponse struct {
	Answer  CardTypesAnswer `xml:"answer"`
	XMLName xml.Name        `xml:"sirena" json:"-"`
}

// CardTypesAnswer is an <answer> section in all card types response
type CardTypesAnswer struct {
	CardTypes CardTypesAnswerDetails `xml:"describe"`
}

// CardTypesAnswerDetails is a <describe> section in all card types response
type CardTypesAnswerDetails struct {
	Data []CardTypesAnswerData `xml:"data"`
}

// CardTypesAnswerData is a <data> section in all card types response
type CardTypesAnswerData struct {
	Code []CardTypesAnswerDataCode `xml:"code"`
	Name []CardTypesAnswerDataName `xml:"name"`
}

// CardTypesAnswerDataCode represents <code> entry in <data> section
type CardTypesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// CardTypesAnswerDataName represents <name> entry in <data> section
type CardTypesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
