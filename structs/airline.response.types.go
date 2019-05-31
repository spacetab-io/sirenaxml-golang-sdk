package structs

import "encoding/xml"

// AirlinesResponse is a response to all airlines request
type AirlinesResponse struct {
	Answer  AirlinesAnswer `xml:"answer"`
	XMLName xml.Name       `xml:"sirena" json:"-"`
}

// AirlinesAnswer is an <answer> section in all airlines response
type AirlinesAnswer struct {
	Airlines AirlinesAnswerDetails `xml:"describe"`
}

// AirlinesAnswerDetails is a <describe> section in all airlines response
type AirlinesAnswerDetails struct {
	Data  []AirlinesAnswerData `xml:"data"`
	Error *Error               `xml:"error"`
}

// AirlinesAnswerData is a <data> section in all airlines response
type AirlinesAnswerData struct {
	AccountCode string                       `xml:"account-code" json:"account_code"`
	Code        []AirlinesAnswerDataCode     `xml:"code" json:"code"`
	Name        []AirlinesAnswerDataName     `xml:"name" json:"name"`
	SubClasses  []AirlinesAnswerDataSubClass `xml:"subclasses>subclass" json:"subclasses,omitempty"`
}

// AirlinesAnswerDataCode represents <lang> entry in <data> section
type AirlinesAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// AirlinesAnswerDataName represents <name> entry in <data> section
type AirlinesAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// AirlinesAnswerDataSubClass represents <subclass> entry in <subclasses> section
type AirlinesAnswerDataSubClass struct {
	Class string `xml:"class,attr" json:"class"`
	Value string `xml:",chardata" json:"value"`
}
