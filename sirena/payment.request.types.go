package sirena

import "encoding/xml"

// PaymentRequest is a <payment-ext-auth> request
type PaymentRequest struct {
	Query   PaymentRequestQuery `xml:"query"`
	XMLName xml.Name            `xml:"sirena"`
}

// PaymentRequestQuery is a <query> section in <payment-ext-auth> request
type PaymentRequestQuery struct {
	Payment PaymentRequestBody `xml:"payment-ext-auth"`
}

// PaymentRequestBody is a body of <payment-ext-auth> request
type PaymentRequestBody struct {
	Regnum  string `xml:"regnum"`
	Surname string `xml:"surname"`
	Action  string `xml:"action"`
	Paydoc  struct {
		Formpay  string `xml:"formpay"`
		Type     string `xml:"type"`
		Num      string `xml:"num,omitempty"`
		ExpDate  string `xml:"exp_date,omitempty"`
		Holder   string `xml:"holder,omitempty"`
		AuthCode string `xml:"auth_code,omitempty"`
	} `xml:"paydoc"`
	Cost *struct {
		Curr  string  `xml:"curr,attr"`
		Value float64 `xml:",chardata"`
	} `xml:"cost,omitempty"`
	RequestParams PaymentRequestParams `xml:"request_params"`
}

// PaymentRequestParams is a <request_params> section in <payment-ext-auth> request
type PaymentRequestParams struct {
	TickSer        string `xml:"tick_ser,omitempty"`
	PaymentTimeout int    `xml:"payment_timeout,omitempty"`
	ReturnReceipt  bool   `xml:"return_receipt,omitempty"`
}
