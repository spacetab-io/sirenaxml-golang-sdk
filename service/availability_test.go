package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func TestService_Avalability(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		logger := logs.NewNullLog()
		sdkClient, err := sdk.NewClient(&sc, logger)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		service := NewSKD(sdkClient)
		checkKeyData(t, sdkClient)
		availabiliteReq := &structs.AvailabilityRequest{
			Query: structs.AvailabilityRequestQuery{
				Availability: structs.Availability{
					Departure: "MOW",
					Arrival:   "LED",
					AnswerParams: structs.AvailabilityAnswerParams{
						ShowFlighttime: true,
					},
				},
			},
		}

		_, err = service.Avalability(availabiliteReq)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
}
