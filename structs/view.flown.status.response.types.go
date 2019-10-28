package structs

import "encoding/xml"

type ViewFlownStatusResponse struct {
	Answer  ViewFlownStatusAnswer `xml:"answer"`
	XMLName xml.Name              `xml:"sirena" json:"-"`
}

type ViewFlownStatusAnswer struct {
	ViewFlownStatus ViewFlownStatusResp `xml:"view_flown_status"`
}

type ViewFlownStatusResp struct {
	Passengers Passenger `xml:"passenger"`
}

type Passenger struct {
	ID      string            `xml:"id,attr"`
	Segment []ViewFlowSegment `xml:"segment"`
}

type ViewFlowSegment struct {
	ID     string `xml:"id,attr"`
	Ticket string `xml:"ticket,attr"`
	EcStat string `xml:"ec_stat,attr"`
}
