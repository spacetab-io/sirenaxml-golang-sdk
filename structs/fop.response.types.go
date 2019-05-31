package structs

import "encoding/xml"

// FOPResponse is a response to all fop request
type FOPResponse struct {
	Answer  FOPAnswer `xml:"answer"`
	XMLName xml.Name  `xml:"sirena" json:"-"`
}

// FOPAnswer is an <answer> section in all fop response
type FOPAnswer struct {
	FOP FOPAnswerDetails `xml:"describe"`
}

// FOPAnswerDetails is a <describe> section in all fop response
type FOPAnswerDetails struct {
	Data  []FOPAnswerData `xml:"data"`
	Error *Error          `xml:"error"`
}

// FOPAnswerData is a <data> section in all fop response
type FOPAnswerData struct {
	Code []FOPAnswerDataCode `xml:"code" json:"code"`
	Name []FOPAnswerDataName `xml:"name" json:"name"`
}

// FOPAnswerDataCode represents <code> entry in <data> section
type FOPAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// FOPAnswerDataName represents <name> entry in <data> section
type FOPAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
