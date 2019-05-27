package structs

import "encoding/xml"

// ErrorResponse is an <error> section in Sirena  response
type ErrorResponse struct {
	Answer  ErrorAnswer `xml:"answer"`
	XMLName xml.Name    `xml:"sirena" json:"-"`
}

type ErrorAnswer struct {
	Error struct {
		Code       int    `xml:"code,attr"`
		CryptError bool   `xml:"crypt_error,attr"`
		Message    string `xml:",chardata"`
	} `xml:"error"`
}
