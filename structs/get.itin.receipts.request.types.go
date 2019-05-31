package structs

import "encoding/xml"

// GetItinReceiptsRequest is a <get_itin_receipts> request
type GetItinReceiptsRequest struct {
	Query   GetItinReceiptsRequestQuery `xml:"query"`
	XMLName xml.Name                    `xml:"sirena"`
}

// GetItinReceiptsRequestQuery is a <query> section in <get_itin_receipts> request
type GetItinReceiptsRequestQuery struct {
	GetItinReceipts GetItinReceipts `xml:"get_itin_receipts"`
}

// GetItinReceipts is a body of <get_itin_receipts> request
type GetItinReceipts struct {
	Regnum  string `xml:"regnum"`
	Surname string `xml:"surname"`
}
