package sdk

import "encoding/xml"

// DocumentsResponse is a response to all documents request
type DocumentsResponse struct {
	Answer  DocumentsAnswer `xml:"answer"`
	XMLName xml.Name        `xml:"sirena" json:"-"`
}

// DocumentsAnswer is an <answer> section in all documents response
type DocumentsAnswer struct {
	Documents DocumentsAnswerDetails `xml:"describe"`
}

// DocumentsAnswerDetails is a <describe> section in all documents response
type DocumentsAnswerDetails struct {
	Data  []DocumentsAnswerData `xml:"data"`
	Error *Error                `xml:"error"`
}

// DocumentsAnswerData is a <data> section in all documents response
type DocumentsAnswerData struct {
	Code             []DocumentsAnswerDataCode `xml:"code" json:"code"`
	Name             []DocumentsAnswerDataName `xml:"name" json:"name"`
	NeedsCitizenship bool                      `xml:"needs_citizenship" json:"needs_citizenship"`
	International    int                       `xml:"international" json:"international"`
	Type             int                       `xml:"type" json:"type"`
}

// DocumentsAnswerDataCode represents <code> entry in <data> section
type DocumentsAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// DocumentsAnswerDataName represents <name> entry in <data> section
type DocumentsAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
