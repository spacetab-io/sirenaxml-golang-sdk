package sirena

import "encoding/xml"

// PaymentTypesRequest is a Sirena <describe> request for all payment types
type PaymentTypesRequest struct {
	Query   PaymentTypesRequestQuery `xml:"query"`
	XMLName xml.Name                 `xml:"sirena"`
}

// PaymentTypesRequestQuery is a <query> section in  all payment types request
type PaymentTypesRequestQuery struct {
	PaymentTypes PaymentTypes `xml:"describe"`
}

// PaymentTypes is a <describe> section in all payment types request
type PaymentTypes struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
