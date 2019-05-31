package structs

import "encoding/xml"

// AddFFInfoResponse is a Sirena response to <add_ff_info> request
type AddFFInfoResponse struct {
	Answer  AddFFInfoAnswer `xml:"answer"`
	XMLName xml.Name        `xml:"sirena" json:"-"`
}

// AddFFInfoAnswer is an <answer> section in Sirena <add_ff_info> response
type AddFFInfoAnswer struct {
	Pult      string                   `xml:"pult,attr,omitempty"`
	AddFFInfo AddFFInfoAnswerAddFFInfo `xml:"add_ff_info"`
}

// AddFFInfoAnswerAddFFInfo is a <add_ff_info> section in Sirena <add_ff_info> response
type AddFFInfoAnswerAddFFInfo struct {
	Ok    *struct{}      `xml:"ok,omitempty"`
	Error *ErrorResponse `xml:"error,omitempty"`
}
