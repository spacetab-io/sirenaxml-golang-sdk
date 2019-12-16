package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var casesGetAssociationPaxInformation = []struct {
	name                 string
	pricingAnswerVariant *PricingAnswerVariant
	result               [][]PricingAnswerVariantFlight
}{
	{
		"common",
		MockPricingAnswerVariant,
		MockDirectionOneFlightGroups,
	},
}

func Test_GetAssociationPaxInformation(t *testing.T) {
	for _, tt := range casesGetAssociationPaxInformation {
		t.Run(tt.name, func(t *testing.T) {
			got := MockPricingAnswerVariantDirectionOne.GetDirectionsFlightGroups(tt.pricingAnswerVariant)

			if !assert.Equal(t, tt.result, got, "should be equal") {
				t.Errorf("GetDirectionsFlightGroups() want: %v, got %v", tt.result, got)
			}
		})
	}
}
