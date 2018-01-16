package sirena

import (
	"encoding/xml"
)

// PricingRouteRequest is a <pricing_route> request
type PricingRouteRequest struct {
	Query   PricingRouteRequestQuery `xml:"query"`
	XMLName xml.Name                 `xml:"sirena"`
}

// PricingRouteRequestQuery is a <query> section in <pricing_route> request
type PricingRouteRequestQuery struct {
	PricingRoute PricingRoute `xml:"pricing_route"`
}

// PricingRoute is a <pricing_route> section
type PricingRoute struct {
	Segments      []PricingRouteRequestSegment   `xml:"segment"`
	Passenger     []PricingRouteRequestPassenger `xml:"passenger"`
	RequestParams PricingRouteRequestParams      `xml:"request_params"`
	AnswerParams  PricingRouteAnswerParams       `xml:"answer_params"`
}

// PricingRouteRequestSegment is a <segment> section in <pricing_route> request
type PricingRouteRequestSegment struct {
	ID          int    `xml:"id,omitempty"`
	JointID     int    `xml:"joint_id,omitempty"`
	Departure   string `xml:"departure"`
	Arrival     string `xml:"arrival"`
	Date        string `xml:"date"`
	Company     string `xml:"company,omitempty"`
	Flight      string `xml:"flight,omitempty"`
	Num         int    `xml:"num,omitempty"`
	Subclass    string `xml:"subclass,omitempty"`
	Class       string `xml:"class,omitempty"`
	Direct      bool   `xml:"direct"`
	Connections string `xml:"connections,omitempty"`
	TimeFrom    int    `xml:"time_from,omitempty"`
	TimeTill    int    `xml:"time_till,omitempty"`
}

// PricingRouteRequestPassenger is a <passenger> section in <pricing_route> request
type PricingRouteRequestPassenger struct {
	Code  string `xml:"code"`
	Count int    `xml:"count"`
	Age   int    `xml:"age,omitempty"`
	Doc   string `xml:"doc,omitempty"`
	Doc2  string `xml:"doc2,omitempty"`
}

// PricingRouteRequestParams is a <request_params> section in <pricing_route> request
type PricingRouteRequestParams struct {
	MinResults        string                      `xml:"min_results,omitempty"`
	MaxResults        string                      `xml:"max_results,omitempty"`
	MixScls           bool                        `xml:"mix_scls,omitempty"`
	MixAc             bool                        `xml:"mix_ac,omitempty"`
	FingeringOrder    string                      `xml:"fingering_order,omitempty"`
	TickSer           string                      `xml:"tick_ser,omitempty"`
	PriceChildAAA     bool                        `xml:"price_child_aaa,omitempty"`
	AsynchronousFares bool                        `xml:"asynchronous_fares,omitempty"`
	Timeout           int                         `xml:"timeout,omitempty"`
	EtIfPossible      bool                        `xml:"et_if_possible,omitempty"`
	ShowVariantTotal  bool                        `xml:"show_varianttotal,omitempty"`
	Formpay           *PricingRouteRequestFormpay `xml:"formpay,omitempty"`
}

// PricingRouteAnswerParams is a <answer_params> section in <pricing_route> request
type PricingRouteAnswerParams struct {
	ShowAvailable    bool   `xml:"show_available,omitempty"`
	ShowIOMatching   bool   `xml:"show_io_matching,omitempty"`
	ShowFlightTime   bool   `xml:"show_flighttime,omitempty"`
	ShowVariantTotal bool   `xml:"show_varianttotal,omitempty"`
	ShowBaseClass    bool   `xml:"show_baseclass,omitempty"`
	ShowRegLatin     bool   `xml:"show_reg_latin,omitempty"`
	ShowUptRec       bool   `xml:"show_upt_rec,omitempty"`
	ShowFareExpDate  bool   `xml:"show_fareexpdate,omitempty"`
	ShowEt           bool   `xml:"show_et,omitempty"`
	ShowNBlanks      bool   `xml:"show_n_blanks,omitempty"`
	Regroup          bool   `xml:"regroup,omitempty"`
	ReturnDate       bool   `xml:"return_date,omitempty"`
	MarkCityPort     bool   `xml:"mark_cityport,omitempty"`
	Lang             string `xml:"lang,omitempty"`
	Curr             string `xml:"curr,omitempty"`
}

// PricingRouteRequestFormpay is a <formpay> element in <request_params>
type PricingRouteRequestFormpay struct {
	Type  string `xml:"type,attr"`
	Num   string `xml:"num,attr"`
	Value string `xml:",chardata"`
}
