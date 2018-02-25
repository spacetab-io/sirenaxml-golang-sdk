package sirena

import "encoding/xml"

// ModifyPNRResponce is a Sirena response to <modify_pnr> request
type ModifyPNRResponce struct {
	Answer  ModifyPNRAnswer `xml:"answer"`
	XMLName xml.Name        `xml:"sirena" json:"-"`
}

// ModifyPNRAnswer is an <answer> section in Sirena <modify_pnr> response
type ModifyPNRAnswer struct {
	Pult      string                   `xml:"pult,attr,omitempty"`
	ModifyPNR ModifyPNRAnswerModifyPNR `xml:"modify_pnr"`
}

// ModifyPNRAnswerModifyPNR is a <modify_pnr> section in Sirena <modify_pnr> response
type ModifyPNRAnswerModifyPNR struct {
	Ok    *struct{} `xml:"ok,omitempty"`
	Error *Error    `xml:"error,omitempty"`
}
