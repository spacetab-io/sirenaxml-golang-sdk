package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r *Repository) Book(
	logAttributes map[string]string,
	bookingSegments []sirena.BookingRequestSegment,
	bookingPassengers []sirena.BookingRequestPassenger,
	bookingContacts *sirena.BookingRequestContacts,
	bookingAgentCommission *sirena.BookingRequestAgentComission,
	brand []sirena.Brand,
) (
	*sirena.BookingResponse,
	error,
) {

	sirenaBookingRequest := sirena.BookingRequest{
		Query: sirena.BookingRequestQuery{
			Booking: sirena.Booking{
				Segments:       bookingSegments,
				Passenger:      bookingPassengers,
				Contacts:       bookingContacts,
				AgentComission: bookingAgentCommission,
				AnswerParams: sirena.BookingAnswerParams{
					ShowUptRec:      true,
					AddRemarks:      true,
					AddRemoteRecloc: true,
					Lang:            "en",
				},
				RequestParams: sirena.BookingRequestParams{
					TickSer:      "",
					ParcelAgency: "",
					Formpay: sirena.BookingRequestFormpay{
						//Type: "CA",
						Value: string(sirena.FormPayCash),
					},

					AllowWaitlist: false,
					Brand:         brand,
				},
			},
		},
	}

	requestBytes, err := xml.MarshalIndent(&sirenaBookingRequest, "  ", "    ")
	if err != nil {
		return nil, err
	}

	responseBytes, err := r.Transport.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var bookingResponse sirena.BookingResponse
	if err = xml.Unmarshal(responseBytes, &bookingResponse); err != nil {
		return nil, err
	}

	return &bookingResponse, nil
}
