package structs

import "encoding/xml"

// ReleaseSeatsResponse is a Sirena response to <release_seats> request
type ReleaseSeatsResponse struct {
	Answer  ReleaseSeatsAnswer `xml:"answer"`
	XMLName xml.Name           `xml:"sirena" json:"-"`
}

// ReleaseSeatsAnswer is an <answer> section in Sirena <release_seats> response
type ReleaseSeatsAnswer struct {
	Pult         string                 `xml:"pult,attr,omitempty"`
	ReleaseSeats ReleaseSeatsAnswerBody `xml:"release_seats"`
}

// ReleaseSeatsAnswerBody is a <release_seats> section in Sirena <release_seats> response
type ReleaseSeatsAnswerBody struct {
	Orders []ReleaseSeatsAnswerOrder `xml:"orders"`
	Error  *ErrorResponse            `xml:"error"`
}

// ReleaseSeatsAnswerOrder is an <order> entry in Sirena <release_seats> response
type ReleaseSeatsAnswerOrder struct {
	Regnum       string `xml:"regnum"`
	BookTime     string `xml:"book_time"`
	Agn          string `xml:"agn,omitempty"`
	PPR          string `xml:"ppr,omitempty"`
	NSeats       int    `xml:"nseats"`
	NSeg         int    `xml:"nseg"`
	NPax         int    `xml:"npax"`
	CommonStatus string `xml:"common_status"`
}
