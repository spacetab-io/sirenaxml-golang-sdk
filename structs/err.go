package structs

import "encoding/xml"

// ErrorResponse receives when connection error occurs
type ErrorResponse struct {
	XMLName xml.Name    `xml:"sirena" json:"-"`
	Answer  ErrorAnswer `xml:"answer"`
}

type ErrorAnswer struct {
	XMLName xml.Name `xml:"answer"`
	Error   Error    `xml:"error"`
}

type Error struct {
	XMLName    xml.Name `xml:"error"`
	Code       int      `xml:"code,attr"`
	CryptError int      `xml:"crypt_error,attr"`
	Message    string   `xml:",chardata"`
}
