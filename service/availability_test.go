package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func TestService_Availability(t *testing.T) {
	tearUp()

	t.Run("success", func(t *testing.T) {
		logger := logs.NewNullLog()

		sdkClient, err := sdk.NewClient(
			logger,
			conf.ClientPrivateKey,
			conf.ClientPrivateKeyPassword,
			conf.ClientPublicKey,
			conf.Ip,
			conf.Environment,
			conf.ServerPublicKey,
			conf.Address,
			conf.Buffer,
			conf.ZippedMessaging,
			conf.MaxConnections,
			conf.ClientID,
		)

		if !assert.NoError(t, err) {
			t.FailNow()
		}

		service := NewSKD(sdkClient)


		checkKeyData(t, sdkClient)
		availabilityRequest := &structs.AvailabilityRequest{
			Query: structs.AvailabilityRequestQuery{
				Availability: structs.Availability{
					Departure: "MOW",
					Arrival:   "LED",
					AnswerParams: structs.AvailabilityAnswerParams{
						ShowFlightTime: true,
					},
				},
			},
		}

		_, err = service.Availability(availabilityRequest)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
}
