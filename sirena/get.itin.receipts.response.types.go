package sirena

import "encoding/xml"

// GetItinReceiptsResponse is a Sirena response to <get_itin_receipts> request
type GetItinReceiptsResponse struct {
	Answer  GetItinReceiptsAnswer `xml:"answer"`
	XMLName xml.Name              `xml:"sirena" json:"-"`
}

// GetItinReceiptsAnswer is an <answer> section in Sirena <get_itin_receipts> response
type GetItinReceiptsAnswer struct {
	Answer          string                        `xml:"answer,attr,omitempty"`
	GetItinReceipts GetItinReceiptsAnswerReceipts `xml:"get_itin_receipts>receipts"`
}

// GetItinReceiptsAnswerReceipts is a <receipts> element in Sirena <get_itin_receipts> response
type GetItinReceiptsAnswerReceipts struct {
	CrTime string `xml:"cr_time,attr"` // "TimeDate" format
	Value  string `xml:",chardata"`
	Error  *Error `xml:"error,omitempty"`
}
