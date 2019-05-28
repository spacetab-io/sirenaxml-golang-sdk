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
	Regnum        string               `xml:"regnum"`
	Surname       string               `xml:"surname"`
	Action        string               `xml:"action"`
	Paydoc        Paydoc               `xml:"paydoc"`
	Cost          *Cost                `xml:"cost,omitempty"`
	RequestParams PaymentRequestParams `xml:"request_params"`
}

// Paydoc is a <paydoc> entry in <payment-ext-auth> request
type Paydoc struct {
	Formpay  string `xml:"formpay"`
	Type     string `xml:"type,omitempty"`
	Num      string `xml:"num,omitempty"`
	ExpDate  string `xml:"exp_date,omitempty"`
	Holder   string `xml:"holder,omitempty"`
	AuthCode string `xml:"auth_code,omitempty"`
}

// Cost is a <cost> entry in <payment-ext-auth> request
type Cost struct {
	Curr  string `xml:"curr,attr"`
	Value string `xml:",chardata"`
}

// PaymentRequestParams is a <request_params> section in <payment-ext-auth> request
type PaymentRequestParams struct {
	TickSer        string `xml:"tick_ser,omitempty"`
	PaymentTimeout int    `xml:"payment_timeout,omitempty"`
	ReturnReceipt  bool   `xml:"return_receipt,omitempty"`
}
