package sirena

import "encoding/xml"

type TaxesResponse struct {
	Answer  TaxesAnswer `xml:"answer"`
	XMLName xml.Name    `xml:"sirena" json:"-"`
}

type TaxesAnswer struct {
	Taxes TaxesAnswerDetails `xml:"describe"`
}

type TaxesAnswerDetails struct {
	Data []TaxesAnswerData `xml:"data"`
}

type TaxesAnswerData struct {
	Code    []TaxesAnswerDataCode    `xml:"code"`
	Name    []TaxesAnswerDataName    `xml:"name"`
	Country []TaxesAnswerDataCountry `xml:"country"`
}

type TaxesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type TaxesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type TaxesAnswerDataCountry struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
