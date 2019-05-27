package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/sdk"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

// AvailabilityXML is a test availability XML
func TestService(t *testing.T) {
	t.Run("test no zipped request", func(t *testing.T) {
		customSirenConfig := sc
		testRequest(t, customSirenConfig)
	})
	t.Run("test zipped request/response", func(t *testing.T) {
		customSirenConfig := sc
		customSirenConfig.ZippedMessaging = true
		testRequest(t, customSirenConfig)
	})
	t.Run("test error params", func(t *testing.T) {
		customSirenConfig := sc
		customSirenConfig.ClientID = 1111
		_, err := sdk.NewClient(&customSirenConfig, lc)
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}

func TestService_Avalability(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		sdkClient, err := sdk.NewClient(&sc, lc)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		service := NewSKD(sdkClient)
		if !assert.NotEmpty(t, sdkClient.Key) {
			t.FailNow()
		}
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

		response, err := service.Avalability(availabiliteReq)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		if !assert.NotEmpty(t, response.Answer.Availability.Flights) {
			t.FailNow()
		}
	})
	//t.Run("error", func(t *testing.T) {
	//	sdkClient, err := sdk.NewClient(&sc, lc)
	//	if !assert.NoError(t, err) {
	//		t.FailNow()
	//	}
	//
	//	service := NewSKD(sdkClient)
	//	if !assert.NotEmpty(t, sdkClient.Key) {
	//		t.FailNow()
	//	}
	//	availabiliteReq := &structs.AvailabilityRequest{
	//		Query: structs.AvailabilityRequestQuery{
	//			Availability: structs.Availability{
	//				Departure: "MOW",
	//				//Arrival:   "LED",
	//				AnswerParams: structs.AvailabilityAnswerParams{
	//					ShowFlighttime: true,
	//				},
	//			},
	//		},
	//	}
	//
	//	response, err := service.Avalability(availabiliteReq)
	//	if !assert.Error(t, err) {
	//		t.FailNow()
	//	}
	//	if !assert.Empty(t, response.Answer.Availability.Flights) {
	//		t.FailNow()
	//	}
	//})
}

func testRequest(t *testing.T, sc configuration.SirenaConfig) {
	sdkClient, err := sdk.NewClient(&sc, lc)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	service := NewSKD(sdkClient)
	if !assert.NotEmpty(t, sdkClient.Key) {
		t.FailNow()
	}
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

	var (
		respChan = make(chan *structs.AvailabilityResponse)
		errChan  = make(chan error)
	)
	for i := 0; i < int(sc.SirenaRequestHandlers); i++ {
		go func() {
			response, err := service.Avalability(availabiliteReq)
			if err != nil {
				errChan <- err
				return
			}
			respChan <- response
		}()
	}

	select {
	case response := <-respChan:
		if !assert.NotEmpty(t, response.Answer.Availability.Flights) {
			t.FailNow()
		}
	case err := <-errChan:
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	}

}
