package sirena

import "encoding/xml"

const (
	Time     = "15:04"
	Date     = "02:01:2006"
	TimeDate = "15:04 02:01:2006"
	DateTime = "02:01:2006 15:04"
)

type BookingResponse struct {
	Answer  BookingAnswer `xml:"answer"`
	XMLName xml.Name      `xml:"sirena" json:"-"`
}

type BookingAnswer struct {
	Pult     string             `xml:"pult,attr,omitempty"`
	MsgID    int                `xml:"msgid,attr"`
	Time     string             `xml:"time,attr"`
	Instance string             `xml:"instance,attr"`
	Booking  BookingAnswerQuery `xml:"booking"`
}

type BookingAnswerQuery struct {
	Regnum   string                `xml:"regnum,attr"`
	Agency   string                `xml:"agency,attr"`
	PNR      BookingAnswerPNR      `xml:"pnr"`
	Contacts BookingAnswerContacts `xml:"contacts"`
	Error    BookingError          `xml:"error"`
}

type BookingError struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:",chardata"`
}

// BookingAnswerPNR is a <PNR> section in <booking> answer
type BookingAnswerPNR struct {
	RegNum            string                      `xml:"regnum"`
	UTCTimeLimit      string                      `xml:"utc_timelimit"`
	TimeLimit         string                      `xml:"timelimit"`
	LatinRegistration bool                        `xml:"latin_registration"`
	Version           int                         `xml:"version"`
	Segments          []BookingAnswerPNRSegment   `xml:"segments"`
	Passengers        []BookingAnswerPNRPassenger `xml:"passengers"`
	Prices            BookingAnswerPNRPrices      `xml:"prices"`
}

// BookingAnswerPNRSegment is a <segment> section in <booking> answer
type BookingAnswerPNRSegment struct {
	ID           int                        `xml:"id,omitempty"`
	JointId      int                        `xml:"joint_id,omitempty"`
	Company      string                     `xml:"company"`
	Flight       string                     `xml:"flight"`
	Subclass     string                     `xml:"subclass"`
	Class        string                     `xml:"class"`
	BaseClass    string                     `xml:"baseclass"`
	SeatCount    int                        `xml:"seatcount"`
	Airplane     string                     `xml:"airplane"`
	Legs         []PNRSegmentLeg            `xml:"legs"`
	Departure    PNRSegmentDepartureArrival `xml:"departure"`
	Arrival      PNRSegmentDepartureArrival `xml:"arrival"`
	Status       PNRSegmentStatus           `xml:"status"`
	FlightTime   string                     `xml:"flightTime"`
	RemoteRecloc string                     `xml:"remote_recloc"`
	Cabin        string                     `xml:"cabin"`
}

type PNRSegmentLeg struct {
	Airplane int                 `xml:"airplane,attr"`
	Dep      PNRSegmentLegDepArr `xml:"dep"`
	Arr      PNRSegmentLegDepArr `xml:"arr"`
}

type PNRSegmentLegDepArr struct {
	TimeLocal string `xml:"time_local,attr"`
	TimeUTC   string `xml:"time_utc,attr"`
	Term      string `xml:"term,attr"`
	Value     string `xml:",chardata"`
}

// PNRSegmentDepartureArrival is a structure for Departure and Arrival fields of <segment> section
type PNRSegmentDepartureArrival struct {
	City     string `xml:"city"`
	Aitport  string `xml:"airport"`
	Date     string `xml:"date"`
	Time     string `xml:"time"`
	Terminal string `xml:"terminal"`
}

type PNRSegmentStatus struct {
	Text   string `xml:"text,attr"`
	Status string `xml:",chardata"`
}

// BookingAnswerPNRPassenger is a <passenger> section in <booking> answer
type BookingAnswerPNRPassenger struct {
	ID          int                  `xml:"id,attr,omitempty"`
	LeadPass    bool                 `xml:"lead_pass,attr"`
	Name        string               `xml:"name"`
	Surname     string               `xml:"surname"`
	Sex         string               `xml:"sex"`
	Birthdate   string               `xml:"birthdate"`
	Age         int                  `xml:"age"`
	DocCode     string               `xml:"doccode"`
	Doc         string               `xml:"doccode"`
	PspExpire   string               `xml:"pspexpire"`
	Category    PNRPassengerCategory `xml:"category"`
	DocCountry  string               `xml:"doc_country"`
	Nationality string               `xml:"nationality"`
	Residence   string               `xml:"residence"`
	Contacts    []Contact            `xml:"contacts"`
}

type PNRPassengerCategory struct {
	RBM      int    `xml:"rbm,attr"`
	Categoty string `xml:",chardata"`
}

// BookingAnswerPNRPrices is a <prices> section in <booking> answer
type BookingAnswerPNRPrices struct {
	TickSer      string                  `xml:"tick_ser,attr"`
	FOP          string                  `xml:"fop,attr"`
	Prices       []BookingAnswerPNRPrice `xml:"prices"`
	VariantTotal PNRVariantTotal         `xml:"variant_total"`
}

type BookingAnswerPNRPrice struct {
	SegmentID         int            `xml:"segment-id,attr"`
	PassengerID       int            `xml:"passenger-id,attr"`
	Code              string         `xml:"code,attr"`
	OrigCode          string         `xml:"orig_code,attr"`
	Count             int            `xml:"count,attr"`
	Currency          string         `xml:"currency,attr"`
	TourCode          string         `xml:"tour_code,attr"`
	FC                string         `xml:"fc,attr"`
	Baggage           string         `xml:"baggage,attr"`
	Ticket            string         `xml:"ticket,attr"`
	ValidatingCompany string         `xml:"validating_company,attr"`
	ACCode            string         `xml:"accode,attr"`
	DocType           string         `xml:"doc_type,attr"`
	DocID             string         `xml:"doc_id,attr"`
	Brand             string         `xml:"brand,attr"`
	Fare              PNRPriceFare   `xml:"fare"`
	Taxes             []PNRPriceTax  `xml:"tax"`
	PaymentInfo       PNRPaymentInfo `xml:"payment_info"`
	Total             float64        `xml:"total"`
}

type PNRPriceFare struct {
	Remark      string           `xml:"remark,attr"`
	FareExpDate string           `xml:"fare_expdate,attr"`
	Value       PNRPriceValue    `xml:"value"`
	Code        PNRPriceFareCode `xml:"code"`
}

type PNRPriceValue struct {
	Value    float64 `xml:",chardata"`
	Currency string  `xml:"currency,attr"`
}

type PNRPriceFareCode struct {
	Code     string `xml:",chardata"`
	BaseCode string `xml:"base_code,attr"`
}

type PNRPriceTax struct {
	Owner string        `xml:"owner,attr"`
	Code  string        `xml:"code"`
	Value PNRPriceValue `xml:"value"`
}

type PNRPaymentInfo struct {
	FOP     string  `xml:"fop,attr"`
	Curr    string  `xml:"curr,attr"`
	Payment float64 `xml:",chardata"`
}

type PNRVariantTotal struct {
	Currency     string  `xml:"currency,attr"`
	VariantTotal float64 `xml:",chardata"`
}

type BookingAnswerContacts struct {
	Contacts []Contact        `xml:"contacts"`
	Customer ContactsCustomer `xml:"customer"`
}

type ContactsCustomer struct {
	FirstName string `xml:"firstname"`
	LastName  string `xml:"lastname"`
}
