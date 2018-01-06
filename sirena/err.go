package sirena

// Error is an <error> section in Sirena  response
type Error struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:",chardata"`
}
