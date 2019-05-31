package structs

import "encoding/xml"

type KeyInfoResponse struct {
	Answer struct {
		KeyInfo struct {
			KeyManager struct {
				ServerPubliKey string `xml:"server_public_key"`
			} `xml:"key_manager"`
		} `xml:"key_info"`
	} `xml:"answer"`
	XMLName xml.Name `xml:"sirena"`
}
