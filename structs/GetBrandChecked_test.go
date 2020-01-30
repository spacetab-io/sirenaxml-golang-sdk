package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesGetBrandChecked = []struct {
	name                          string
	paxType                       string
	pricingAnswerVariantDirection *PricingAnswerVariant
	result                        bool
}{
	{
		"passed variant without brand type",
		"ADT",
		MockPricingAnswerVariant,
		false,
	},
	{
		"passed variant with brand",
		"ADT",
		MockPricingAnswerVariantWithBrand,
		true,
	},
	{
		"passed variant with brand, and bad pax type",
		"CHD",
		MockPricingAnswerVariantWithBrand,
		false,
	},
}

func Test_GetBrandChecked(t *testing.T) {
	for _, tt := range casesGetBrandChecked {
		t.Run(tt.name, func(t *testing.T) {
			got := MockPricingAnswerVariantFlightGroup.GetBrandChecked(*tt.pricingAnswerVariantDirection, tt.paxType)

			if !assert.Equal(t, tt.result, got, "should be equal") {
				t.Errorf("GetBrandChecked() want%v, got %v", tt.result, got)
			}
		})
	}
}
