package structs

import "encoding/xml"

// GetItinReceiptsDataRequest is a <get_itin_receipts_data> request
type GetItinReceiptsDataRequest struct {
	Query   GetItinReceiptsDataRequestQuery `xml:"query"`
	XMLName xml.Name                        `xml:"sirena"`
}

// GetItinReceiptsDataRequestQuery is a <query> section in <get_itin_receipts_data> request
type GetItinReceiptsDataRequestQuery struct {
	GetItinReceiptsData GetItinReceiptsData `xml:"get_itin_receipts_data"`
}

// GetItinReceiptsData is a body of <get_itin_receipts_data> request
type GetItinReceiptsData struct {
	Regnum  string `xml:"regnum"`
	Surname string `xml:"surname"`
}
