package structs

import "encoding/xml"

// AddRemarkResponse is a Sirena response to <add_remark> request
type AddRemarkResponse struct {
	Answer  AddRemarkAnswer `xml:"answer"`
	XMLName xml.Name        `xml:"sirena" json:"-"`
}

// AddRemarkAnswer is an <answer> section in Sirena <add_remark> response
type AddRemarkAnswer struct {
	Pult      string                   `xml:"pult,attr,omitempty"`
	AddRemark AddRemarkAnswerAddRemark `xml:"add_remark"`
}

// AddRemarkAnswerAddRemark is a <add_remark> section in Sirena <add_remark> response
type AddRemarkAnswerAddRemark struct {
	Ok    *struct{}      `xml:"ok,omitempty"`
	Error *ErrorResponse `xml:"error,omitempty"`
}
