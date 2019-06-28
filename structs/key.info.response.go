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
	Key             KeyData    `xml:"key"`
	Expiration      SirenaTime `xml:"expiration,omitempty"`
	ServerPublicKey string     `xml:"server_public_key"`
}

type KeyData struct {
	State string `xml:"state,attr"`
	Key   string `xml:",chardata"`
}

type SirenaTime struct {
	time.Time
}

func (c *SirenaTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	parse, err := time.Parse(TimeDate, v)
	if err != nil {
		return err
	}
	*c = SirenaTime{parse}
	return nil
}
