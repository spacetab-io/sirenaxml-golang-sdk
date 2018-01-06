package sirena

type PricingResponse struct {
	Answer PricingAnswer `xml:"answer"`
}

type PricingAnswer struct {
	Pricing          *PricingAnswerPricing `xml:"pricing,omitempty"`
	PricingMonobrand *PricingAnswerPricing `xml:"pricing_mono_brand,omitempty"`
}

type PricingAnswerPricing struct {
	Variants []PricingAnswerVariant `xml:"variant"`
	Flights  []*PricingAnswerFlight `xml:"flight"`
	Error    *Error                 `xml:"error"`
}

type PricingAnswerFlight struct {
	ID          int                   `xml:"id,attr"`
	Company     string                `xml:"company"`
	Num         int                   `xml:"num"`
	Flight      string                `xml:"flight"`
	Origin      PricingAnswerLocation `xml:"origin"`
	Destination PricingAnswerLocation `xml:"destination"`
	DeptDate    string                `xml:"deptdate"`
	ArrvDate    string                `xml:"arrvdate"`
	DeptTime    string                `xml:"depttime"`
	ArrvTime    string                `xml:"arrvtime"`
	Airplane    string                `xml:"airplane"`
}

type PricingAnswerVariant struct {
	FlightGroups []PricingAnswerVariantFlightGroup `xml:"flights"`
	Directions   []PricingAnswerVariantDirection   `xml:"direction"`
	Total        PricingAnswerVariantTotal         `xml:"variant_total"`
}

type PricingAnswerVariantTotal struct {
	Currency string  `xml:"currency,attr"`
	Total    float64 `xml:",chardata"`
}

type PricingAnswerVariantFlightGroup struct {
	Flight []PricingAnswerVariantFlight `xml:"flight"`
}

type PricingAnswerVariantFlight struct {
	ID        int    `xml:"id,attr"`
	Num       int    `xml:"num,attr"`
	SubClass  string `xml:"subclass,attr"`
	BaseClass string `xml:"baseclass,attr"`
	Available int    `xml:"available,attr"`
}

type PricingAnswerVariantDirection struct {
	Num    int                   `xml:"num,attr"`
	Prices []*PricingAnswerPrice `xml:"price"`
}

type PricingAnswerLocation struct {
	City     string `xml:"city,attr"`
	Terminal string `xml:"terminal,attr"`
	Value    string `xml:",chardata"`
}

type PricingAnswerPrice struct {
	// PassengerID       int                     `xml:"passenger-id,attr"`
	// Code              string                  `xml:"code,attr"`
	// Count             int                     `xml:"count,attr"`
	// Currency          string                  `xml:"currency,attr"`
	// Ticket            string                  `xml:"ticket,attr"`
	// FC                string                  `xml:"fc,attr"`
	// DocID             string                  `xml:"doc_id,attr"`
	// ACCode            string                  `xml:"accode,attr"`
	// ValidatingCompany string                  `xml:"validating_company,attr"`
	// FOP               string                  `xml:"fop,attr"`
	OrigCode string `xml:"orig_code,attr"`
	// OrigID            int                     `xml:"orig_id,attr"`
	Fare  *PricingAnswerPriceFare `xml:"fare"`
	Taxes []PricingAnswerPriceTax `xml:"tax"`
	// Total             float64                 `xml:"total"`
}

type PricingAnswerPriceFare struct {
	Remark      string  `xml:"remark,attr"`
	FareExpDate string  `xml:"fare_expdate,attr"`
	Code        string  `xml:"code,attr"`
	BaseCode    string  `xml:"base_code,attr"`
	Total       float64 `xml:",chardata"`
}

type PricingAnswerPriceTax struct {
	Code  string  `xml:"code,attr"`
	Owner string  `xml:"owner,attr"`
	Total float64 `xml:",chardata"`
}
