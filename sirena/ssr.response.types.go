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
	Data []SSRAnswerData `xml:"data"`
}

// SSRAnswerData is a <data> section in all special requests response
type SSRAnswerData struct {
	Code              []SSRAnswerDataCode `xml:"code"`
	Name              []SSRAnswerDataName `xml:"name"`
	Category          int                 `xml:"category"`
	ForSegment        int                 `xml:"for_segment"`
	ForPassenger      int                 `xml:"for_passenger"`
	FreeText          int                 `xml:"free_text"`
	TravelPortAllowed bool                `xml:"travelport_allowed"`
}

// SSRAnswerDataCode represents <code> entry in <data> section
type SSRAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// SSRAnswerDataName represents <name> entry in <data> section
type SSRAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
