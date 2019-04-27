package sdk

import "encoding/xml"

// BookingRequest is a Sirena booking request
type BookingRequest struct {
	Query   BookingRequestQuery `xml:"query"`
	XMLName xml.Name            `xml:"sirena"`
}

// BookingRequestQuery is a <query> section in Sirena booking request
type BookingRequestQuery struct {
	Booking Booking `xml:"booking"`
}

// Booking is a <booking> section in Sirena booking request
type Booking struct {
	Segments        []BookingRequestSegment        `xml:"segment"`
	LandSegments    []BookingRequestLandSegment    `xml:"land_segment"`
	StandbySegments []BookingRequestStandbySegment `xml:"standby_segment"`
	Passenger       []BookingRequestPassenger      `xml:"passenger"`
	Contacts        *BookingRequestContacts        `xml:"contacts,omitempty"`
	SpecialServices *BookingRequestSpecialServices `xml:"special_services,omitempty"`
	Remarks         *BookingRequestRemarks         `xml:"remarks,omitempty"`
	AgentComission  *BookingRequestAgentComission  `xml:"agent_comission,omitempty"`
	AnswerParams    BookingAnswerParams            `xml:"answer_params,omitempty"`
	RequestParams   BookingRequestParams           `xml:"request_params,omitempty"`
}

// BookingRequestParams is a <request_params> section in Sirena booking request
type BookingRequestParams struct {
	TickSer       string                `xml:"tick_ser,omitempty"`
	ParcelAgency  string                `xml:"parcel_agency,omitempty"`
	Formpay       BookingRequestFormpay `xml:"formpay,omitempty"`
	AllowWaitlist bool                  `xml:"allow_waitlist,omitempty"`
	Brand         string                `xml:"brand,omitempty"`
}

// BookingRequestFormpay is a <formpay> entry in <request_params>
type BookingRequestFormpay struct {
	Type  string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// BookingAnswerParams is an <answer_params> entry in Sirena booking request
type BookingAnswerParams struct {
	ShowUptRec      bool   `xml:"show_upt_rec,omitempty"`
	AddRemarks      bool   `xml:"add_remarks,omitempty"`
	AddSsr          bool   `xml:"add_ssr,omitempty"`
	AddRemoteRecloc bool   `xml:"add_remote_recloc,omitempty"`
	ShowComission   bool   `xml:"show_comission,omitempty"`
	Lang            string `xml:"lang,omitempty"`
}

// BookingRequestAgentComission is an <agent_comission> in Sirena booking request
type BookingRequestAgentComission struct {
	Type  string `xml:"type,attr"`
	Curr  string `xml:"curr,attr,omitempty"`
	Value string `xml:",chardata"`
}

// BookingRequestRemarks is a <remarks> section in Sirena booking request
type BookingRequestRemarks struct {
	Remark []string `xml:"remark,omitempty"`
}

// BookingRequestSpecialServices is a <special_services> section in Sirena booking request
type BookingRequestSpecialServices struct {
	Ssrs []Ssr `xml:"ssr"`
}

type Ssr struct {
	Type   string `xml:"type,attr"`
	Text   string `xml:"text,attr,omitempty"`
	PassId int    `xml:"pass_id,attr,omitempty"`
	SegId  int    `xml:"seg_id,attr,omitempty"`
}

type BookingRequestContacts struct {
	Phone *Phone `xml:"phone"`
	Email string `xml:"email"`
}

type BookingRequestPassenger struct {
	ID          int       `xml:"-"`
	Lastname    string    `xml:"lastname"`
	Firstname   string    `xml:"firstname"`
	Category    string    `xml:"category"`
	Sex         string    `xml:"sex"`
	Birthdate   string    `xml:"birthdate"`
	Doccode     string    `xml:"doccode"`
	Doc         string    `xml:"doc"`
	PspExpire   string    `xml:"pspexpire,omitempty"`
	DocCountry  string    `xml:"doc_country,omitempty"`
	Nationality string    `xml:"nationality,omitempty"`
	Residence   string    `xml:"residence,omitempty"`
	DocCodeDisc string    `xml:"doccode_disc,omitempty"`
	DocDisc     string    `xml:"doc_disc,omitempty"`
	Phones      []Phone   `xml:"phone,omitempty"`
	Contacts    []Contact `xml:"contact,omitempty"`
}

type Phone struct {
	ContID  int    `xml:"cont_id,attr"`
	LocID   int    `xml:"loc_id,attr"`
	Type    string `xml:"type,attr,omitempty"`
	Comment string `xml:"comment,attr,omitempty"`
	Value   string `xml:",chardata"`
}

// type Email struct {
// 	ContID int    `xml:"cont_id,attr"`
// 	LocID  int    `xml:"loc_id,attr"`
// 	Type   string `xml:"type,attr"`
// 	Email  string `xml:",chardata"`
// }

type Contact struct {
	ContID  int    `xml:"cont_id,attr"`
	LocID   int    `xml:"loc_id,attr"`
	Type    string `xml:"type,attr,omitempty"`
	Comment string `xml:"comment,attr,omitempty"`
	Value   string `xml:",chardata"`
}

type BookingRequestStandbySegment struct {
	Company   string `xml:"company"`
	Flight    string `xml:"flight,omitempty"`
	Departure string `xml:"departure"`
	DepTime   string `xml:"depTime,omitempty"`
	Arrival   string `xml:"arrival"`
	ArrTime   string `xml:"arrTime,omitempty"`
	Date      string `xml:"date,omitempty"`
	Subclass  string `xml:"subclass"`
	ID        int    `xml:"id,omitempty"`
	JointId   int    `xml:"joint_id,omitempty"`
}

type BookingRequestLandSegment struct {
	ID      int `xml:"id,omitempty"`
	JointId int `xml:"joint_id,omitempty"`
}

// BookingRequestSegment is a <segment> section in <booking> request
type BookingRequestSegment struct {
	Company   string `xml:"company"`
	Flight    string `xml:"flight"`
	Departure string `xml:"departure"`
	Arrival   string `xml:"arrival"`
	Date      string `xml:"date"`
	Subclass  string `xml:"subclass"`
	ID        int    `xml:"id,omitempty"`
	JointId   int    `xml:"joint_id,omitempty"`
}
