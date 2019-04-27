package sdk

import "encoding/xml"

// GetItinReceiptsResponse is a Sirena response to <get_itin_receipts> request
type GetItinReceiptsResponse struct {
	Answer  GetItinReceiptsAnswer `xml:"answer"`
	XMLName xml.Name              `xml:"sirena" json:"-"`
}

// GetItinReceiptsAnswer is an <answer> section in Sirena <get_itin_receipts> response
type GetItinReceiptsAnswer struct {
	Answer          string              `xml:"answer,attr,omitempty"`
	GetItinReceipts GetItinReceiptsBody `xml:"get_itin_receipts"`
}

// GetItinReceiptsBody is a body of <get_itin_receipts> response
type GetItinReceiptsBody struct {
	Receipts *GetItinReceiptsAnswerReceipts `xml:"receipts"`
	Error    *Error                         `xml:"error,omitempty"`
}

// GetItinReceiptsAnswerReceipts is a <receipts> element in Sirena <get_itin_receipts> response
type GetItinReceiptsAnswerReceipts struct {
	CrTime string `xml:"cr_time,attr"` // "TimeDate" format
	Value  string `xml:",chardata"`
}
