package sdk

import "encoding/xml"

// DocumentsRequest is a Sirena <describe> request for all documents
type DocumentsRequest struct {
	Query   DocumentsRequestQuery `xml:"query"`
	XMLName xml.Name              `xml:"sirena"`
}

// DocumentsRequestQuery is a <query> section in  all documents request
type DocumentsRequestQuery struct {
	Documents Documents `xml:"describe"`
}

// Documents is a <describe> section in all documents request
type Documents struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
