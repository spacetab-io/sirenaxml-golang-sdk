package structs

import "encoding/xml"

// TaxesResponse is a response to all taxes request
type TaxesResponse struct {
	Answer  TaxesAnswer `xml:"answer"`
	XMLName xml.Name    `xml:"sirena" json:"-"`
}

// TaxesAnswer is an <answer> section in all taxes response
type TaxesAnswer struct {
	Taxes TaxesAnswerDetails `xml:"describe"`
}

// TaxesAnswerDetails is a <describe> section in all taxes response
type TaxesAnswerDetails struct {
	Data  []TaxesAnswerData `xml:"data"`
	Error *Error            `xml:"error"`
}

// TaxesAnswerData is a <data> section in all taxes response
type TaxesAnswerData struct {
	Code []TaxesAnswerDataCode `xml:"code" json:"code"`
	Name []TaxesAnswerDataName `xml:"name" json:"name"`
}

// TaxesAnswerDataCode represents <code> entry in <data> section
type TaxesAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// TaxesAnswerDataName represents <name> entry in <data> section
type TaxesAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// TaxesAnswerDataCountry represents <country> entry in <data> section
type TaxesAnswerDataCountry struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
