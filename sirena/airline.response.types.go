package sirena

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
	Data []AirlinesAnswerData `xml:"data"`
}

// AirlinesAnswerData is a <data> section in all airlines response
type AirlinesAnswerData struct {
	AccountCode string                       `xml:"account-code"`
	Code        []AirlinesAnswerDataCode     `xml:"code"`
	Name        []AirlinesAnswerDataName     `xml:"name"`
	SubClasses  []AirlinesAnswerDataSubClass `xml:"subclasses>subclass"`
}

// AirlinesAnswerDataCode represents <lang> entry in <data> section
type AirlinesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// AirlinesAnswerDataName represents <name> entry in <data> section
type AirlinesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// AirlinesAnswerDataSubClass represents <subclass> entry in <subclasses> section
type AirlinesAnswerDataSubClass struct {
	Class string `xml:"class,attr"`
	Value string `xml:",chardata"`
}
