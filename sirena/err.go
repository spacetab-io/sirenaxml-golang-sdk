package sirena

// ErrorResponse is an <error> section in Sirena  response
type ErrorResponse struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:",chardata"`
}
