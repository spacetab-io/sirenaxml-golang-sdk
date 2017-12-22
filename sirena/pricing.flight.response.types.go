package sirena

type PricingFlightResponse struct {
	Answer PricingFlightAnswer `xml:"answer"`
}

type PricingFlightAnswer struct {
	Pricing PricingAnswerPricing `xml:"pricing_flight"`
}
