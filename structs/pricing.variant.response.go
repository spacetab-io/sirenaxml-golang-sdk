package structs

type PricingVariantResponse struct {
	Answer PricingRouteAnswer `xml:"answer"`
}

type PricingVariantAnswer struct {
	PricingRoute *PricingRouteAnswerPricingRoute `xml:"pricing_route,omitempty"`
}

type PricingVariantAnswerPricingRoute struct {
	BrandInfo []BrandInfo                 `xml:"brand_info"`
	Variants  []PricingVariantAnswerVariant `xml:"variant"`
	Flights   []*PricingVariantAnswerFlight `xml:"flight"`
	Error     *Error                      `xml:"error,omitempty"`
}


type PricingVariantAnswerFlight struct {
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

type PricingVariantAnswerVariant struct {
	Flight []PricingVariantAnswerVariantFlight `xml:"flight"`
	Total  PricingVariantAnswerVariantTotal    `xml:"variant_total"`

	//Directions   []PricingRouteAnswerVariantDirection   `xml:"direction"`
}

type PricingVariantAnswerVariantTotal struct {
	Currency string  `xml:"currency,attr"`
	Total    float64 `xml:",chardata"`
}

type PricingVariantAnswerVariantFlightGroup struct {
	Flight []PricingVariantAnswerVariantFlight `xml:"flight"`
}

type PricingVariantAnswerVariantFlight struct {
	Price     Price  `xml:"price"`
	ID        int    `xml:"id,attr"`
	Num       int    `xml:"num,attr"`
	SubClass  string `xml:"subclass,attr"`
	BaseClass string `xml:"baseclass,attr"`
	Available int    `xml:"available,attr"`
}

type PricingVariantAnswerVariantDirection struct {
	Num    int                        `xml:"num,attr"`
	Prices []*PricingVariantAnswerPrice `xml:"price"`
}

type PricingVariantAnswerLocation struct {
	City     string `xml:"city,attr"`
	Terminal string `xml:"terminal,attr"`
	Value    string `xml:",chardata"`
}

type PricingVariantAnswerPrice struct {
	Brand    string                       `xml:"brand"`
	OrigCode string                       `xml:"orig_code,attr"`
	Fare     *PricingVariantAnswerPriceFare `xml:"fare"`
	Taxes    []PricingVariantAnswerPriceTax `xml:"tax"`
}

type PricingVariantAnswerPriceFare struct {
	Remark      string  `xml:"remark,attr"`
	FareExpDate string  `xml:"fare_expdate,attr"`
	Code        string  `xml:"code,attr"`
	BaseCode    string  `xml:"base_code,attr"`
	Total       float64 `xml:",chardata"`
}

type PricingVariantAnswerPriceTax struct {
	Code  string  `xml:"code,attr"`
	Owner string  `xml:"owner,attr"`
	Total float64 `xml:",chardata"`
}
