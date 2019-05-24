package sirena

import "encoding/xml"

// FOPRequest is a Sirena <describe> request for all fop (формы оплаты)
type FOPRequest struct {
	Query   FOPRequestQuery `xml:"query"`
	XMLName xml.Name        `xml:"sirena"`
}

// FOPRequestQuery is a <query> section in all fop request
type FOPRequestQuery struct {
	FOP FOP `xml:"describe"`
}

// FOP is a <describe> section in all fop request
type FOP struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
