package sdk

import "encoding/xml"

// OrderCancelResponse is a Sirena response to <booking-cancel> request
type OrderCancelResponse struct {
	Answer  OrderCancelAnswer `xml:"answer"`
	XMLName xml.Name          `xml:"sirena" json:"-"`
}

// OrderCancelAnswer is an <answer> section in Sirena order cancel response
type OrderCancelAnswer struct {
	Cancel OrderCancelData `xml:"booking-cancel"`
}

// OrderCancelData is a <booking-cancel> entry in Sirena booking cancel response
type OrderCancelData struct {
	OK    *struct{}      `xml:"ok,omitempty"`
	Error *ErrorResponse `xml:"error,omitempty"`
}
