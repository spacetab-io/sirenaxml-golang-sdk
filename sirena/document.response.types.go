package sirena

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
	Data []DocumentsAnswerData `xml:"data"`
}

// DocumentsAnswerData is a <data> section in all documents response
type DocumentsAnswerData struct {
	Code             []DocumentsAnswerDataCode `xml:"code"`
	Name             []DocumentsAnswerDataName `xml:"name"`
	NeedsCitizenship bool                      `xml:"needs_citizenship"`
	International    int                       `xml:"international"`
	Type             int                       `xml:"type"`
}

// DocumentsAnswerDataCode represents <code> entry in <data> section
type DocumentsAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// DocumentsAnswerDataName represents <name> entry in <data> section
type DocumentsAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
