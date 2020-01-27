package proxy

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func TestStorage_GetAvailability(t *testing.T) {
	nl := logs.NewNullLog()
	p := MockPublisher{}
	proxyStorage := NewStorage(p, proxyPath, nl, false)
	req := &structs.AvailabilityRequest{
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
	reqXML, err := xml.Marshal(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	resp, err := proxyStorage.GetAvailability(reqXML)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	if !assert.NotNil(t, resp) {
		t.FailNow()
	}

	if resp.Answer.Availability.Flights == nil && resp.Answer.Availability.Flight == nil {
		assert.NotEmpty(t, resp.Answer.Availability.Flights)
		assert.NotEmpty(t, resp.Answer.Availability.Flight)
	}

}
