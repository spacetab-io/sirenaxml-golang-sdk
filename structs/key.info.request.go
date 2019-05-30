package structs

import "encoding/xml"

// KeyInfoRequest
type KeyInfoRequest struct {
	Query struct {
		KeyInfo struct{} `xml:"key_info"`
	} `xml:"query"`
	XMLName xml.Name `xml:"sirena"`
}
