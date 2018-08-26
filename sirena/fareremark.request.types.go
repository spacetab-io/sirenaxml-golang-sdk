package sirena

import "encoding/xml"

// FareRemarkRequest is a <fareremark> request
type FareRemarkRequest struct {
	Query   FareRemarkRequestQuery `xml:"query"`
	XMLName xml.Name               `xml:"sirena"`
}

// FareRemarkRequestQuery is a <query> section in <fareremark> request
type FareRemarkRequestQuery struct {
	FareRemark FareRemark `xml:"fareremark"`
}

// FareRemark is a body of <fareremark> request
type FareRemark struct {
	Company       string                  `xml:"company,omitempty"`
	Code          string                  `xml:"code"`
	RequestParams FareRemarkRequestParams `xml:"request_params"`
	AnswerParams  FareRemarkAnswerParams  `xml:"answer_params"`
}

// FareRemarkRequestParams is a <request_params> section in <fareremark> request
type FareRemarkRequestParams struct {
	Cat int      `xml:"cat,omitempty"`
	Upt PriceUpt `xml:"upt"`
}

// FareRemarkAnswerParams is a <answer_params> section in <fareremark> request
type FareRemarkAnswerParams struct {
	Lang string `xml:"lang,omitempty"`
}
