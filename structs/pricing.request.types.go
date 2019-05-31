package structs

import (
	"encoding/xml"
)

// PricingRequest is a <pricing> request
type PricingRequest struct {
	Query   PricingRequestQuery `xml:"query"`
	XMLName xml.Name            `xml:"sirena"`
}

// PricingFlightRequest is a <pricing> request
type PricingFlightRequest struct {
	Query   PricingFlightRequestQuery `xml:"query"`
	XMLName xml.Name                  `xml:"sirena"`
}

// PricingRequestQuery is a <query> section in <pricing> request
type PricingRequestQuery struct {
	Pricing Pricing `xml:"pricing"`
}

// PricingFlightRequestQuery is a <query> section in <pricing> request
type PricingFlightRequestQuery struct {
	Pricing Pricing `xml:"pricing_flight"`
}

type Pricing struct {
	Segments      []PricingRequestSegment   `xml:"segment"`
	Passenger     []PricingRequestPassenger `xml:"passenger"`
	RequestParams PricingRequestParams      `xml:"request_params"`
	AnswerParams  PricingAnswerParams       `xml:"answer_params"`
}

// PricingRequestSegment is a <segment> section in <pricing> request
type PricingRequestSegment struct {
	ID          int                            `xml:"id,omitempty"`
	JointID     int                            `xml:"joint_id,omitempty"`
	Departure   string                         `xml:"departure"`
	Arrival     string                         `xml:"arrival"`
	Date        string                         `xml:"date"`
	Company     string                         `xml:"company,omitempty"`
	Flight      string                         `xml:"flight,omitempty"`
	Num         string                         `xml:"num,omitempty"`
	Subclass    string                         `xml:"subclass,omitempty"`
	Class       []string                       `xml:"class,omitempty"`
	Direct      bool                           `xml:"direct"`
	Connections string                         `xml:"connections,omitempty"`
	TimeFrom    int                            `xml:"time_from,omitempty"`
	TimeTill    int                            `xml:"time_till,omitempty"`
	Ignore      []PricingRequestIgnoredAirline `xml:"ignore>acomp,omitempty"`
}

// PricingRequestPassenger is a <passenger> section in <pricing> request
type PricingRequestPassenger struct {
	Code  string `xml:"code"`
	Count int    `xml:"count"`
	Age   int    `xml:"age,omitempty"`
	Doc   string `xml:"doc,omitempty"`
	Doc2  string `xml:"doc2,omitempty"`
}

// PricingRequestParams is a <request_params> section in <pricing> request
type PricingRequestParams struct {
	MinResults        string                 `xml:"min_results,omitempty"`
	MaxResults        int                    `xml:"max_results,omitempty"`
	MixScls           bool                   `xml:"mix_scls,omitempty"`
	MixAc             bool                   `xml:"mix_ac,omitempty"`
	FingeringOrder    string                 `xml:"fingering_order,omitempty"`
	TickSer           string                 `xml:"tick_ser,omitempty"`
	PriceChildAAA     bool                   `xml:"price_child_aaa,omitempty"`
	AsynchronousFares bool                   `xml:"asynchronous_fares,omitempty"`
	Timeout           int                    `xml:"timeout,omitempty"`
	EtIfPossible      bool                   `xml:"et_if_possible,omitempty"`
	Formpay           *PricingRequestFormpay `xml:"formpay,omitempty"`
	// ShowVariantTotal  bool                   `xml:"show_varianttotal,omitempty"`
}

// PricingAnswerParams is a <answer_params> section in <pricing> request
type PricingAnswerParams struct {
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

// PricingRequestFormpay is a <formpay> element in <request_params>
type PricingRequestFormpay struct {
	Type  string `xml:"type,attr"`
	Num   string `xml:"num,attr"`
	Value string `xml:",chardata"`
}

// PricingRequestIgnoredAirline is an <acomp> element in <pricing> request
type PricingRequestIgnoredAirline struct {
	Name   string `xml:"name,attr"`
	Flight string `xml:"flight"`
}
