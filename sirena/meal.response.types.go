package sirena

import "encoding/xml"

// MealsResponse is a response to all meals request
type MealsResponse struct {
	Answer  MealsAnswer `xml:"answer"`
	XMLName xml.Name    `xml:"sirena" json:"-"`
}

// MealsAnswer is an <answer> section in all meals response
type MealsAnswer struct {
	Meals MealsAnswerDetails `xml:"describe"`
}

// MealsAnswerDetails is a <describe> section in all meals response
type MealsAnswerDetails struct {
	Data  []MealsAnswerData `xml:"data"`
	Error *Error            `xml:"error"`
}

// MealsAnswerData is a <data> section in all meals response
type MealsAnswerData struct {
	Code []MealsAnswerDataCode `xml:"code" json:"code"`
	Name []MealsAnswerDataName `xml:"name" json:"name"`
}

// MealsAnswerDataCode represents <code> entry in <data> section
type MealsAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// MealsAnswerDataName represents <name> entry in <data> section
type MealsAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
