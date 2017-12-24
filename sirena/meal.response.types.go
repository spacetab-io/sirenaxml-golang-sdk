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
	Data []MealsAnswerData `xml:"data"`
}

// MealsAnswerData is a <data> section in all meals response
type MealsAnswerData struct {
	Code    []MealsAnswerDataCode    `xml:"code"`
	Name    []MealsAnswerDataName    `xml:"name"`
	Country []MealsAnswerDataCountry `xml:"country"`
}

// MealsAnswerDataCode represents <code> entry in <data> section
type MealsAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// MealsAnswerDataName represents <name> entry in <data> section
type MealsAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// MealsAnswerDataCountry represents <country> entry in <data> section
type MealsAnswerDataCountry struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
