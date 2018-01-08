package sirena

import "encoding/xml"

// ReleaseSeatsRequest is a <release_seats> request
type ReleaseSeatsRequest struct {
	Query   ReleaseSeatsRequestQuery `xml:"query"`
	XMLName xml.Name                 `xml:"sirena"`
}

// ReleaseSeatsRequestQuery is a <query> section in <release_seats> request
type ReleaseSeatsRequestQuery struct {
	ReleaseSeats ReleaseSeatsRequestBody `xml:"release_seats"`
}

// ReleaseSeatsRequestBody is a body of <release_seats> request
type ReleaseSeatsRequestBody struct {
	Regnum    string                       `xml:"regnum"`
	Surname   string                       `xml:"surname"`
	Passenger ReleaseSeatsRequestPassenger `xml:"passenger"`
}

// ReleaseSeatsRequestPassenger is a <passenger> entry in <release_seats> request
type ReleaseSeatsRequestPassenger struct {
	Name    string `xml:"name"`
	Surname string `xml:"surname"`
}
