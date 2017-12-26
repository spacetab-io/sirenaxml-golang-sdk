package sirena

import "encoding/xml"

type OrderResponse struct {
	Answer  OrderAnswer `xml:"answer"`
	XMLName xml.Name    `xml:"sirena" json:"-"`
}

type OrderAnswer struct {
	Pult     string           `xml:"pult,attr,omitempty"`
	MsgID    int              `xml:"msgid,attr"`
	Time     string           `xml:"time,attr"`
	Instance string           `xml:"instance,attr"`
	Order    OrderAnswerQuery `xml:"order"`
}

type OrderAnswerQuery struct {
	Regnum   string              `xml:"regnum,attr"`
	Agency   string              `xml:"agency,attr"`
	PNR      BookingAnswerPNR    `xml:"pnr"`
	Tickinfo OrderAnswerTickinfo `xml:"tickinfo"`
}

// OrderAnswerPNR is a <PNR> section in <order> answer
type OrderAnswerPNR struct {
	Bdate              string                      `xml:"bdate, attr"`
	Fop                string                      `xml:"fop, attr"`
	Curr               string                      `xml:"curr, attr"`
	Sum                float64                     `xml:"sum, attr"`
	RegNum             string                      `xml:"regnum"`
	UTCTimeLimit       string                      `xml:"utc_timelimit"`
	TimeLimit          string                      `xml:"timelimit"`
	LatinRegistration  bool                        `xml:"latin_registration"`
	Version            int                         `xml:"version"`
	Contacts           BookingAnswerContacts       `xml:"contacts"`
	Segments           []BookingAnswerPNRSegment   `xml:"segments"`
	Passengers         []BookingAnswerPNRPassenger `xml:"passengers"`
	Prices             BookingAnswerPNRPrices      `xml:"prices"`
	CommonStatus       string                      `xml:"common_status"`
	PossibleActionList PNRPossibleActionList       `xml:"possible_action_list"`
}

type PNRPossibleActionList struct {
	Action string `xml:"action"`
}

type OrderAnswerTickinfo struct {
	Ticknum string `xml:"ticknum,attr"`
	SegID   int    `xml:"seg_id,attr"`
	PassID  int    `xml:"pass_id,attr"`
	Value   string `xml:",chardata"`
}
