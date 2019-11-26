package structs

import "encoding/xml"

// PNRStatusResponse is a Sirena response to <pnr_status> request
type PNRStatusResponse struct {
	Answer  PNRStatusAnswer `xml:"answer"`
	XMLName xml.Name        `xml:"sirena" json:"-"`
}

// PNRStatusAnswer is an <answer> section in Sirena <pnr_status> response
type PNRStatusAnswer struct {
	Pult      string                   `xml:"pult,attr,omitempty"`
	MsgID     int                      `xml:"msgid,attr,omitempty"`
	Time      string                   `xml:"time,attr,omitempty"` //TimeSecDate format
	Instance  string                   `xml:"instance,attr,omitempty"`
	PNRStatus PNRStatusAnswerPNRStatus `xml:"pnr_status"`
}

// PNRStatusAnswerPNRStatus is a <pnr_status> section in Sirena <pnr_status> response
type PNRStatusAnswerPNRStatus struct {
	Regnum           string             `xml:"regnum"`
	Agn              string             `xml:"agn,omitempty"`
	Ppr              string             `xml:"ppr,omitempty"`
	NSeats           int                `xml:"nseats,omitempty"`
	NSeg             int                `xml:"nseg,omitempty"`
	NPax             int                `xml:"npax,omitempty"`
	TimeLimit        string             `xml:"timelimit,omitempty"`          // DateTime format
	UTCTimeLimit     string             `xml:"utc_timelimit,omitempty"`      // TimeDate format
	BookTime         string             `xml:"book_time,omitempty"`          // DateTime format
	VoidTimeLimitUTC string             `xml:"void_timelimit_utc,omitempty"` // TimeDate format
	Segments         []PNRStatusSegment `xml:"segments>segment"`
	Tickinfo         PNRStatusTickinfo  `xml:"tickinfo,omitempty"`
	NewTickinfo      []PNRStatusTicket  `xml:"new_tickinfo>ticket"`
	CommonStatus     string             `xml:"common_status"`
	Error            *Error             `xml:"error,omitempty"`
}

// PNRStatusSegment is a <segment> subsection of <segments> section in Sirena <pnr_status> response
type PNRStatusSegment struct {
	SegID     int    `xml:"seg_id,attr,omitempty"`
	NSeats    int    `xml:"nseats,attr,omitempty"`
	BookTime  string `xml:"book_time,attr,omitempty"` // DateTime format
	CreatedBy int    `xml:"created_by,attr,omitempty"`
}

// PNRStatusTickinfo is a <tickinfo> section in Sirena <pnr_status> response
type PNRStatusTickinfo struct {
	Ticknum   string `xml:"ticknum,attr,omitempty"`
	TickSer   string `xml:"tick_ser,attr,omitempty"`
	IsEtick   bool   `xml:"is_etick,attr,omitempty"`
	AcCode    string `xml:"accode,attr,omitempty"`
	TktPpr    string `xml:"tkt_ppr,attr,omitempty"`
	PrintTime string `xml:"print_time,attr,omitempty"` // TimeDate format
	SegID     int    `xml:"seg_id,attr,omitempty"`
	PassID    int    `xml:"pass_id,attr,omitempty"`
	Value     string `xml:",chardata"`
}

// PNRStatusTicket is a <ticket> subsection of <new_tickinfo> section in Sirena <pnr_status> response
type PNRStatusTicket struct {
	Ser     string                `xml:"ser,attr,omitempty"`
	Num     string                `xml:"num,attr,omitempty"`
	IsEtick bool                  `xml:"is_etick,attr,omitempty"`
	Mco     bool                  `xml:"mco,attr,omitempty"`
	Action  PNRStatusTicketAction `xml:"action"`
}

// PNRStatusTicketAction is an <action> element in <ticket> subsection of <new_tickinfo> section in Sirena <pnr_status> response
type PNRStatusTicketAction struct {
	Type        string  `xml:"type,attr,omitempty"`
	Ontime      string  `xml:"ontime,attr,omitempty"` // "DateTime" format
	NSeats      int     `xml:"n_seats,attr,omitempty"`
	SegID       int     `xml:"seg_id,attr,omitempty"`
	Opr         string  `xml:"opr,attr,omitempty"`
	Pult        string  `xml:"pult,attr,omitempty"`
	Sum         float64 `xml:"sum,attr,omitempty"`
	Curr        string  `xml:"curr,attr,omitempty"`
	OrigSer     string  `xml:"orig_ser,attr,omitempty"`
	OrigTicknum string  `xml:"orig_ticknum,attr,omitempty"`
	OrigSegID   int     `xml:"orig_seg_id,attr,omitempty"`
}
