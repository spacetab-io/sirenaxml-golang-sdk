package structs

type PricingRouteResponse struct {
	Answer PricingRouteAnswer `xml:"answer"`
}

type PricingRouteAnswer struct {
	PricingRoute *PricingRouteAnswerPricingRoute `xml:"pricing_route,omitempty"`
}

type PricingRouteAnswerPricingRoute struct {
	Variants []PricingRouteAnswerVariant `xml:"variant"`
	Flights  []*PricingRouteAnswerFlight `xml:"flight"`
	Error    *Error                      `xml:"error,omitempty"`
}

type Price struct {
	Brand             string `xml:"brand,attr"`
	PassengerID       string `xml:"passenger-id,attr"`
	Code              string `xml:"code,attr"`
	Count             string `xml:"count,attr"`
	Currency          string `xml:"currency,attr"`
	Ticket            string `xml:"ticket,attr"`
	TicketCpn         string `xml:"ticket_cpn,attr"`
	Baggage           string `xml:"baggage,attr"`
	Fc                string `xml:"fc,attr"`
	DocType           string `xml:"doc_type,attr"`
	DocID             string `xml:"doc_id,attr"`
	Accode            string `xml:"accode,attr"`
	ValidatingCompany string `xml:"validating_company,attr"`
	Fop               string `xml:"fop,attr"`
	OrigCode          string `xml:"orig_code,attr"`
	OrigID            string `xml:"orig_id,attr"`
	Fare              struct {
		Remark      string `xml:"remark,attr"`
		FareExpdate string `xml:"fare_expdate,attr"`
		Code        string `xml:"code,attr"`
		BaseCode    string `xml:"base_code,attr"`
	} `xml:"fare"`
	Total string `xml:"total"`
}

type PricingRouteAnswerFlight struct {
	Price       Price                      `xml:"price"`
	ID          int                        `xml:"id,attr"`
	Company     string                     `xml:"company"`
	Num         int                        `xml:"num"`
	Flight      string                     `xml:"flight"`
	Origin      PricingRouteAnswerLocation `xml:"origin"`
	Destination PricingRouteAnswerLocation `xml:"destination"`
	DeptDate    string                     `xml:"deptdate"`
	ArrvDate    string                     `xml:"arrvdate"`
	DeptTime    string                     `xml:"depttime"`
	ArrvTime    string                     `xml:"arrvtime"`
	Airplane    string                     `xml:"airplane"`
	FlightTime  string                     `xml:"flightTime"`
}

type PricingRouteAnswerVariant struct {
	Flight []PricingRouteAnswerVariantFlight `xml:"flight"`
	Total  PricingRouteAnswerVariantTotal    `xml:"variant_total"`

	//Directions   []PricingRouteAnswerVariantDirection   `xml:"direction"`
}

type PricingRouteAnswerVariantTotal struct {
	Currency string  `xml:"currency,attr"`
	Total    float64 `xml:",chardata"`
}

type PricingRouteAnswerVariantFlightGroup struct {
	Flight []PricingRouteAnswerVariantFlight `xml:"flight"`
}

type PricingRouteAnswerVariantFlight struct {
	Price     Price  `xml:"price"`
	ID        int    `xml:"id,attr"`
	Num       int    `xml:"num,attr"`
	SubClass  string `xml:"subclass,attr"`
	BaseClass string `xml:"baseclass,attr"`
	Available int    `xml:"available,attr"`
}

type PricingRouteAnswerVariantDirection struct {
	Num    int                        `xml:"num,attr"`
	Prices []*PricingRouteAnswerPrice `xml:"price"`
}

type PricingRouteAnswerLocation struct {
	City     string `xml:"city,attr"`
	Terminal string `xml:"terminal,attr"`
	Value    string `xml:",chardata"`
}

type PricingRouteAnswerPrice struct {
	Brand    string                       `xml:"brand"`
	OrigCode string                       `xml:"orig_code,attr"`
	Fare     *PricingRouteAnswerPriceFare `xml:"fare"`
	Taxes    []PricingRouteAnswerPriceTax `xml:"tax"`
}

type PricingRouteAnswerPriceFare struct {
	Remark      string  `xml:"remark,attr"`
	FareExpDate string  `xml:"fare_expdate,attr"`
	Code        string  `xml:"code,attr"`
	BaseCode    string  `xml:"base_code,attr"`
	Total       float64 `xml:",chardata"`
}

type PricingRouteAnswerPriceTax struct {
	Code  string  `xml:"code,attr"`
	Owner string  `xml:"owner,attr"`
	Total float64 `xml:",chardata"`
}
