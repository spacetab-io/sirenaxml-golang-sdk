package sirena

import "encoding/xml"

// TaxesRequest is a <describe> request
type TaxesRequest struct {
	Query   TaxesRequestQuery `xml:"query"`
	XMLName xml.Name          `xml:"sirena"`
}

// TaxesRequestQuery is a <query> section in <describe> request
type TaxesRequestQuery struct {
	Taxes Taxes `xml:"describe"`
}

type Taxes struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
