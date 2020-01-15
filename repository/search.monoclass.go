package repository

import (
	"encoding/xml"
	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (s Repository) SearchMonoclass(logAttributes map[string]string, pricingSegments []sirena.PricingRequestSegment, pricingPassengers []sirena.PricingRequestPassenger, maxResults, timeout int) (*sirena.PricingResponse, error) {

	pricingRequest := sirena.PricingRequest{
		Query: sirena.PricingRequestQuery{
			Pricing: sirena.Pricing{
				Segments:  pricingSegments,
				Passenger: pricingPassengers,
				RequestParams: sirena.PricingRequestParams{
					Formpay: &sirena.PricingRequestFormpay{
						Value: sirena.FormPayCash,
					},
					MinResults:     "spOnePass",
					MaxResults:     maxResults,
					MixScls:        true,
					MixAc:          true,
					FingeringOrder: "differentFlightsCombFirst", //"differentFlightsFirst",
					//TickSer:           "ЭБМ",
					PriceChildAAA:     true,
					AsynchronousFares: false,
					EtIfPossible:      true,
					ShowBaggageInfo:   true,
					Timeout:           timeout,
				},
				AnswerParams: sirena.PricingAnswerParams{
					ShowAvailable:    true,
					ShowIOMatching:   true,
					ShowFlightTime:   true,
					ShowVariantTotal: true,
					ShowBaseClass:    true,
					ShowRegLatin:     true,
					ShowUptRec:       true,
					ShowFareExpDate:  true,
					ShowEt:           true,
					ShowNBlanks:      true,
					Regroup:          true,
					ReturnDate:       true,
					MarkCityPort:     true,
					ShowUPT18Cat:     true,
					ShowTimeLimit:    true,
					Lang:             "en",
					Curr:             "rub",
				},
			},
		},
	}

	requestBytes, err := xml.MarshalIndent(&pricingRequest, "  ", "    ")
	if err != nil {
		return nil, err
	}

	//log.Printf("Search request: \n %s", string(requestBytes))

	pricingResponseXML, err := s.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	//log.Printf("Search response: \n %s", string(pricingResponseXML))


	var pricingResponse sirena.PricingResponse
	if err := xml.Unmarshal(pricingResponseXML, &pricingResponse); err != nil {
		return nil, err
	}

	return &pricingResponse, nil
}
