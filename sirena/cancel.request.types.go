package sirena

import "encoding/xml"

// CancelRequest is a <payment-ext-auth:refund> request
type CancelRequest struct {
	Query   CancelRequestQuery `xml:"query"`
	XMLName xml.Name           `xml:"sirena"`
}

// CancelRequestQuery is a <query> section in <payment-ext-auth:refund> request
type CancelRequestQuery struct {
	Cancel CancelRequestBody `xml:"booking-cancel"`
}

// CancelRequestBody is a body of <payment-ext-auth:refund> request
type CancelRequestBody struct {
	Regnum  string `xml:"regnum"`
	Surname string `xml:"surname"`
}
