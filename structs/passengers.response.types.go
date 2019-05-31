package structs

import "encoding/xml"

// PassengersResponse is a response to all Passengers request
type PassengersResponse struct {
	Answer  PassengersAnswer `xml:"answer"`
	XMLName xml.Name         `xml:"sirena" json:"-"`
}

// PassengersAnswer is an <answer> section in all passenger(s) response
type PassengersAnswer struct {
	Passengers PassengersAnswerDetails `xml:"describe"`
}

// PassengersAnswerDetails is a <describe> section in passenger(s) response
type PassengersAnswerDetails struct {
	Data  []PassengersAnswerData `xml:"data"`
	Error *Error                 `xml:"error"`
}

// PassengersAnswerData is a <data> section in all passenger(s) response
type PassengersAnswerData struct {
	Code          []PassengersAnswerDataCode `xml:"code" json:"code"`
	Name          []PassengersAnswerDataName `xml:"name" json:"name"`
	AgeCategory   int                        `xml:"agecat" json:"age_category"` // 0 - ADT, 1 - CHD, 2 - INF
	GroupCategory int                        `xml:"group" json:"group_category"`
	Seats         int                        `xml:"seats" json:"seats"`            // Number of seats for category
	SpecCategory  int                        `xml:"spec_cat" json:"spec_category"` // Special category tag
}

// PassengersAnswerDataCode represents <code> entry in <data> section
type PassengersAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// PassengersAnswerDataName represents <name> entry in <data> section
type PassengersAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// PassengersAnswerDataCountry represents <country> entry in <data> section
type PassengersAnswerDataCountry struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
