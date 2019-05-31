package structs

import "encoding/xml"

// RegionsRequest is a <describe> request
type RegionsRequest struct {
	Query   RegionsRequestQuery `xml:"query"`
	XMLName xml.Name            `xml:"sirena"`
}

// RegionsRequestQuery is a <query> section in <describe> request
type RegionsRequestQuery struct {
	Regions Regions `xml:"describe"`
}

// Regions is a <describe> section in all regions request
type Regions struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
