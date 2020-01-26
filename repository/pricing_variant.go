package repository

//
//func (s Storage) PricingVariant(requestID string, passengers []sirena.PricingVariantRequestPassenger, segments []sirena.PricingVariantRequestSegment) (*sirena.PricingVariantResponse, error) {
//
//	answerParems := sirena.PricingVariantAnswerParams{
//		ShowBrandInfo: true,
//	}
//
//	query := sirena.PricingVariantRequestQuery{sirena.PricingVariant{
//		segments,
//		passengers,
//		sirena.PricingVariantRequestParams{},
//		answerParems,
//	}}
//
//	pricingRouteRequest := sirena.PricingVariantRequest{
//		Query:   query,
//		XMLName: xml.Name{},
//	}
//
//	// Encode sirena booking request into XML
//	requestBytes, err := xml.MarshalIndent(&pricingRouteRequest, "  ", "    ")
//	if err != nil {
//		return nil, err
//	}
//	// Debug Sirena request
//	//logs.Log.Debugf("Request:\n%s", requestBytes)
//
//	logAttributes := publisher.NewLogAttributes(
//		requestID,
//		logging.FlowStepBook,
//		"booking",
//	)
//
//	responseBytes, err := s.Request(requestBytes, logAttributes)
//	if err != nil {
//		return nil, err
//	}
//	//logs.Log.Debugf("response:\n%s", responseBytes)
//
//	var pricingRouteResponseResponse sirena.PricingVariantResponse
//	if err = xml.Unmarshal(responseBytes, &pricingRouteResponseResponse); err != nil {
//		return nil, err
//	}
//
//	//sirenaPricingRouteError := pricingRouteResponseResponse.Answer.PricingRoute.Error
//
//	return &pricingRouteResponseResponse, nil
//}
