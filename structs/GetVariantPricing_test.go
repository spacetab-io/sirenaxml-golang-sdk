package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var casesGetVariantPricing = []struct {
	name                 string
	paxType              string
	pricingAnswerVariant *PricingAnswerVariant
	result               *PricingAnswerPrice
}{
	{
		"passed ADT pax type",
		"ADT",
		MockPricingAnswerVariant,
		&MockPricingAnswerPrice,
	},
	{
		"passed inexistent pax type",
		"CHD",
		MockPricingAnswerVariant,
		nil,
	},
}

func Test_GetVariantPricing(t *testing.T) {
	for _, tt := range casesGetVariantPricing {
		t.Run(tt.name, func(t *testing.T) {
			got := MockPricingAnswerVariantFlightOne.GetVariantPricing(*tt.pricingAnswerVariant, tt.paxType)

			if !assert.Equal(t, tt.result, got, "should be equal") {
				t.Errorf("GetVariantPricing(%v) want:%v, got %v", tt.paxType, tt.pricingAnswerVariant, got)
			}
		})
	}
}
