package sirena

import "encoding/xml"

// PCardTypesResponse is a response to all pcard_types request
type PCardTypesResponse struct {
	Answer  PCardTypesAnswer `xml:"answer"`
	XMLName xml.Name         `xml:"sirena" json:"-"`
}

// PCardTypesAnswer is an <answer> section in all pcard_types response
type PCardTypesAnswer struct {
	PCardTypes PCardTypesAnswerDetails `xml:"describe"`
}

// PCardTypesAnswerDetails is a <describe> section in all pcard_types response
type PCardTypesAnswerDetails struct {
	Data  []PCardTypesAnswerData `xml:"data"`
	Error *Error                 `xml:"error"`
}

// PCardTypesAnswerData is a <data> section in all pcard_types response
type PCardTypesAnswerData struct {
	Code []PCardTypesAnswerDataCode `xml:"code" json:"code"`
	Name []PCardTypesAnswerDataName `xml:"name" json:"name"`
}

// PCardTypesAnswerDataCode represents <code> entry in <data> section
type PCardTypesAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// PCardTypesAnswerDataName represents <name> entry in <data> section
type PCardTypesAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
