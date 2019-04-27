package sirena

import "encoding/xml"

// OrderResponse is a Sirena response to <order> request
type OrderResponse struct {
	Answer  OrderAnswer `xml:"answer"`
	XMLName xml.Name    `xml:"sirena" json:"-"`
}

// OrderAnswer is an <answer> section in Sirena order response
type OrderAnswer struct {
	Pult     string           `xml:"pult,attr,omitempty"`
	MsgID    int              `xml:"msgid,attr"`
	Time     string           `xml:"time,attr"`
	Instance string           `xml:"instance,attr"`
	Order    OrderAnswerOrder `xml:"order"`
}

// OrderAnswerOrder is an <order> section in Sirena <order> response
type OrderAnswerOrder struct {
	Regnum   string              `xml:"regnum,attr"`
	Agency   string              `xml:"agency,attr"`
	PNR      OrderAnswerPNR      `xml:"pnr"`
	Tickinfo OrderAnswerTickinfo `xml:"tickinfo"`
	Error    *ErrorResponse      `xml:"error"`
}

// OrderAnswerPNR is a <pnr> section in Sirena order response
type OrderAnswerPNR struct {
	Bdate              string                      `xml:"bdate,attr"`
	Fop                string                      `xml:"fop,attr"`
	Curr               string                      `xml:"curr,attr"`
	Sum                float64                     `xml:"sum,attr"`
	Regnum             string                      `xml:"regnum"`
	UTCTimeLimit       string                      `xml:"utc_timelimit"`
	TimeLimit          string                      `xml:"timelimit"`
	LatinRegistration  bool                        `xml:"latin_registration"`
	Version            int                         `xml:"version"`
	Contacts           BookingAnswerContacts       `xml:"contacts"`
	Segments           []BookingAnswerPNRSegment   `xml:"segments>segment"`
	Passengers         []BookingAnswerPNRPassenger `xml:"passengers>passenger"`
	Prices             BookingAnswerPNRPrices      `xml:"prices"`
	CommonStatus       string                      `xml:"common_status"`
	PossibleActionList PNRPossibleActionList       `xml:"possible_action_list"`
}

// PNRPossibleActionList is a <possible_action_list> entry in <pnr> section
type PNRPossibleActionList struct {
	Action string `xml:"action"`
}

// OrderAnswerTickinfo is a <tickinfo> entry in <order> section
type OrderAnswerTickinfo struct {
	Ticknum string `xml:"ticknum,attr"`
	SegID   int    `xml:"seg_id,attr"`
	PassID  int    `xml:"pass_id,attr"`
	Value   string `xml:",chardata"`
}
