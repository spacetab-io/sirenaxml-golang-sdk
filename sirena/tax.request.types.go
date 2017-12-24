package sirena

import "encoding/xml"

// TaxesRequest is a Sirena <describe> request for all taxes
type TaxesRequest struct {
	Query   TaxesRequestQuery `xml:"query"`
	XMLName xml.Name          `xml:"sirena"`
}

// TaxesRequestQuery is a <query> section in  all taxes request
type TaxesRequestQuery struct {
	Taxes Taxes `xml:"describe"`
}

// Taxes is a <describe> section in all taxes request
type Taxes struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
