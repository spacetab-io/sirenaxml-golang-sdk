package structs

import "encoding/xml"

// ErrorResponse is an <error> section in Sirena  response
type ErrorResponse struct {
	Answer struct {
		Error Error `xml:"error"`
	} `xml:"answer"`
	XMLName xml.Name `xml:"sirena" json:"-"`
}

type Error struct {
	Code       int    `xml:"code,attr"`
	CryptError bool   `xml:"crypt_error,attr"`
	Message    string `xml:",chardata"`
}
