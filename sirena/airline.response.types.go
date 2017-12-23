package sirena

import "encoding/xml"

type AirlinesResponse struct {
	Answer  AirlinesAnswer `xml:"answer"`
	XMLName xml.Name       `xml:"sirena"`
}

type AirlinesAnswer struct {
	Airlines AirlinesAnswerDetails `xml:"describe"`
	Data     string                `xml:"data,attr"`
}

type AirlinesAnswerDetails struct {
	Data []AirlinesAnswerData `xml:"data"`
}

type AirlinesAnswerData struct {
	AccountCode string                       `xml:"account-code"`
	Code        []AirlinesAnswerDataCode     `xml:"code"`
	Name        []AirlinesAnswerDataName     `xml:"name"`
	SubClasses  []AirlinesAnswerDataSubClass `xml:"subclasses>subclass"`
}

type AirlinesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type AirlinesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type AirlinesAnswerDataSubClass struct {
	Class string `xml:"class,attr"`
	Value string `xml:",chardata"`
}
