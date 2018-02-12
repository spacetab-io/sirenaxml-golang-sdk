package sirena

// PricingFlightResponse is a Sirena response to <pricing_flight> request
type PricingFlightResponse struct {
	Answer PricingFlightAnswer `xml:"answer"`
}

// PricingFlightAnswer is an <answer> entry in <pricing_flight> response
type PricingFlightAnswer struct {
	Pricing *PricingAnswerPricing `xml:"pricing_flight"`
}
