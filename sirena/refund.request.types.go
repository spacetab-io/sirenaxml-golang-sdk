package sirena

import "encoding/xml"

// RefundRequest is a <payment-ext-auth:refund> request
type RefundRequest struct {
	Query   RefundRequestQuery `xml:"query"`
	XMLName xml.Name           `xml:"sirena"`
}

// RefundRequestQuery is a <query> section in <payment-ext-auth:refund> request
type RefundRequestQuery struct {
	Refund RefundRequestBody `xml:"payment-ext-auth"`
}

// RefundRequestBody is a body of <payment-ext-auth:refund> request
type RefundRequestBody struct {
	Regnum        string                    `xml:"regnum"`
	Surname       string                    `xml:"surname"`
	Action        string                    `xml:"action"`
	Mode          string                    `xml:"mode,omitempty"`
	RequestParams RefundRequestParams       `xml:"request_params,omitempty"`
	AnswerParams  RefundRequestAnswerParams `xml:"answer_params,omitempty"`
}

// RefundRequestParams is a <request_params> section in <payment-ext-auth:refund> request
type RefundRequestParams struct {
	Pretend bool `xml:"pretend,omitempty"`
}

// RefundRequestAnswerParams is a <answer_params> section in <payment-ext-auth:refund> request
type RefundRequestAnswerParams struct {
	ShowPaydoc       bool `xml:"show_paydoc,omitempty"`
	ShowPriceDetails bool `xml:"show_price_details,omitempty"`
}
