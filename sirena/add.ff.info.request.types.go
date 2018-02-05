package sirena

import "encoding/xml"

// AddFFInfoRequest is a <add_ff_info> request
type AddFFInfoRequest struct {
	Query   AddFFInfoRequestQuery `xml:"query"`
	XMLName xml.Name              `xml:"sirena"`
}

// AddFFInfoRequestQuery is a <query> section in <add_ff_info> request
type AddFFInfoRequestQuery struct {
	AddFFInfo AddFFInfo `xml:"add_ff_info"`
}

// AddFFInfo is a body of <add_ff_info> request
type AddFFInfo struct {
	Regnum    string                      `xml:"regnum"`
	Surname   string                      `xml:"surname"`
	Passenger []AddFFInfoRequestPassenger `xml:"passenger"`
	ID        int                         `xml:"id,attr,omitempty"`
}

// AddFFInfoRequestPassenger is a passenger in <add_ff_info> request
type AddFFInfoRequestPassenger struct {
	FreqFlierID PassengerFreqFlierID               `xml:"freq_flier_id"`
	Segment     []AddFFInfoRequestPassengerSegment `xml:"segment"`
}

// AddFFInfoRequestPassengerSegment is a <segment> subsection of <passenger> section of <add_ff_info> request
type AddFFInfoRequestPassengerSegment struct {
	ID int `xml:"id,attr,omitempty"`
}

// PassengerFreqFlierID is a <freq_flier_id> subsection of <passenger> section of <add_ff_info> request
type PassengerFreqFlierID struct {
	IssuedBy string `xml:"issued_by,attr,omitempty"`
	Value    uint   `xml:",chardata"`
}
