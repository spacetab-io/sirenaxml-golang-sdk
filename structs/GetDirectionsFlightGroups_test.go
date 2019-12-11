package structs

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var MockPricingAnswerVariant = &PricingAnswerVariant{
	FlightGroups: []PricingAnswerVariantFlightGroup{{
		Flight: []PricingAnswerVariantFlight{{
			ID:         "1",
			Num:        1,
			Class:      "Y",
			SubClass:   "Y",
			BaseClass:  "B",
			Available:  0,
			SegmentNum: "",
			Cabin:      "N",
		}, {
			ID:         "5",
			Num:        1,
			Class:      "f",
			SubClass:   "q",
			BaseClass:  "B",
			Available:  0,
			SegmentNum: "",
			Cabin:      "N",
		}, {
			ID:         "2",
			Num:        2,
			Class:      "Z",
			SubClass:   "H",
			BaseClass:  "J",
			Available:  0,
			SegmentNum: "",
			Cabin:      "J",
		}, {
			ID:         "3",
			Num:        2,
			Class:      "Z",
			SubClass:   "H",
			BaseClass:  "J",
			Available:  0,
			SegmentNum: "",
			Cabin:      "J",
		}},
	}},
	Directions: []PricingAnswerVariantDirection{MockPricingAnswerVariantDirectionOne, MockPricingAnswerVariantDirectionTwo},
	Total:      PricingAnswerVariantTotal{},
}

var MockPricingAnswerVariantDirectionOne = PricingAnswerVariantDirection{
	Num: 1,
	Prices: []*PricingAnswerPrice{{
		Upt:               PriceUpt{},
		Fare:              nil,
		Taxes:             nil,
		Vat:               nil,
		Baggage:           "",
		ValidatingAirline: "",
		OriginalPaxType:   "",
		BrandCode:         "",
		Total:             0,
		Currency:          "",
		PassengerID:       0,
		PaxType:           "",
	}},
}

var MockPricingAnswerVariantDirectionTwo = PricingAnswerVariantDirection{
	Num: 2,
	Prices: []*PricingAnswerPrice{{
		Upt:               PriceUpt{},
		Fare:              nil,
		Taxes:             nil,
		Vat:               nil,
		Baggage:           "",
		ValidatingAirline: "",
		OriginalPaxType:   "",
		BrandCode:         "",
		Total:             0,
		Currency:          "",
		PassengerID:       0,
		PaxType:           "",
	}},
}

var casesGetAssociationPaxInformation = []struct {
	name                          string
	paxType                       string
	pricingAnswerVariantDirection *PricingAnswerVariant
}{
	{
		"passed ADT pax type",
		"ADT",
		MockPricingAnswerVariant,
	},
	{
		"passed inexistent pax type",
		"CHD",
		MockPricingAnswerVariant,
	},
}

func Test_GetAssociationPaxInformation(t *testing.T) {
	for _, tt := range casesGetAssociationPaxInformation {
		t.Run(tt.name, func(t *testing.T) {
			got := MockPricingAnswerVariantDirectionOne.GetDirectionsFlightGroups(tt.pricingAnswerVariantDirection)

			spew.Dump(got)

			//if !assert.Equal(t, got, "should be equal") {
			//	t.Errorf("getPaxByID(%v) got %v", tt.paxType, got)
			//}
		})
	}
}
