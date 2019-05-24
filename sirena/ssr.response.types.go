package sirena

import "encoding/xml"

// SSRResponse is a response to all special requests request
type SSRResponse struct {
	Answer  SSRAnswer `xml:"answer"`
	XMLName xml.Name  `xml:"sirena" json:"-"`
}

// SSRAnswer is an <answer> section in all special requests response
type SSRAnswer struct {
	SSR SSRAnswerDetails `xml:"describe"`
}

// SSRAnswerDetails is a <describe> section in all special requests response
type SSRAnswerDetails struct {
	Data  []SSRAnswerData `xml:"data"`
	Error *Error          `xml:"error"`
}

// SSRAnswerData is a <data> section in all special requests response
type SSRAnswerData struct {
	Code              []SSRAnswerDataCode `xml:"code" json:"code"`
	Name              []SSRAnswerDataName `xml:"name" json:"name"`
	Category          int                 `xml:"category" json:"category"`
	ForSegment        int                 `xml:"for_segment" json:"for_segment"`
	ForPassenger      int                 `xml:"for_passenger" json:"for_passenger"`
	FreeText          int                 `xml:"free_text" json:"free_text"`
	TravelPortAllowed bool                `xml:"travelport_allowed" json:"travelport_allowed"`
}

// SSRAnswerDataCode represents <code> entry in <data> section
type SSRAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// SSRAnswerDataName represents <name> entry in <data> section
type SSRAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
