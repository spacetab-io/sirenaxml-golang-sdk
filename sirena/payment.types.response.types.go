package sirena

import "encoding/xml"

// PaymentTypesResponse is a response to all payment types request
type PaymentTypesResponse struct {
	Answer  PaymentTypesAnswer `xml:"answer"`
	XMLName xml.Name           `xml:"sirena" json:"-"`
}

// PaymentTypesAnswer is an <answer> section in all payment types response
type PaymentTypesAnswer struct {
	PaymentTypes PaymentTypesAnswerDetails `xml:"describe"`
}

// PaymentTypesAnswerDetails is a <describe> section in all payment types response
type PaymentTypesAnswerDetails struct {
	Data []PaymentTypesAnswerData `xml:"data"`
}

// PaymentTypesAnswerData is a <data> section in all payment types response
type PaymentTypesAnswerData struct {
	Code []PaymentTypesAnswerDataCode `xml:"code"`
	Name []PaymentTypesAnswerDataName `xml:"name"`
}

// PaymentTypesAnswerDataCode represents <code> entry in <data> section
type PaymentTypesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// PaymentTypesAnswerDataName represents <name> entry in <data> section
type PaymentTypesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
