package structs

import (
	"encoding/xml"
	"time"
)

type KeyInfoResponse struct {
	Answer struct {
		KeyInfo struct {
			KeyManager KeyManager `xml:"key_manager"`
		} `xml:"key_info"`
	} `xml:"answer"`
	XMLName xml.Name `xml:"sirena"`
}

type KeyManager struct {
	Key             KeyData   `xml:"key"`
	Expiration      time.Time `xml:"expiration"`
	ServerPublicKey string    `xml:"server_public_key"`
}

type KeyData struct {
	State string `xml:"state,attr"`
	Key   string `xml:",chardata"`
}
