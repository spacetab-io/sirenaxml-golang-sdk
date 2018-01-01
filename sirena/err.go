package sirena

// SirenaError is an <error> section in Sirena  response
type SirenaError struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:",chardata"`
}
