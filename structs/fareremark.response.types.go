package structs

import "encoding/xml"

// FareRemarkResponse is a Sirena response to <fareremark> request
type FareRemarkResponse struct {
	Answer  FareRemarkAnswer `xml:"answer"`
	XMLName xml.Name         `xml:"sirena" json:"-"`
}

// FareRemarkAnswer is an <answer> section in Sirena <fareremark> response
type FareRemarkAnswer struct {
	Pult       string                     `xml:"pult,attr,omitempty"`
	FareRemark FareRemarkAnswerFareRemark `xml:"fareremark"`
}

// FareRemarkAnswerFareRemark is a <fareremark> section in Sirena <fareremark> response
type FareRemarkAnswerFareRemark struct {
	NewFare bool           `xml:"new_fare,omitempty"`
	Remark  string         `xml:"remark"`
	Error   *Error `xml:"error,omitempty"`
}
