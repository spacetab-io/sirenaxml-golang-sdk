package sirena

import "encoding/xml"

// PNRStatusRequest is a <pnr_status> request
type PNRStatusRequest struct {
	Query   PNRStatusQuery `xml:"query"`
	XMLName xml.Name       `xml:"sirena"`
}

// PNRStatusQuery is a <query> section in <pnr_status> request
type PNRStatusQuery struct {
	PNRStatus PNRStatus `xml:"pnr_status"`
}

// PNRStatus is a body of <pnr_status> request
type PNRStatus struct {
	Regnum       string                `xml:"regnum"`
	AnswerParams PNRStatusAnswerParams `xml:"answer_params"`
}

// PNRStatusAnswerParams is a <answer_params> section in <pnr_status> request
type PNRStatusAnswerParams struct {
	Tickinfo        bool `xml:"tickinfo,omitempty"`
	AddCommonStatus bool `xml:"add_common_status,omitempty"`
	MoreInfo        bool `xml:"more_info,omitempty"`
}
