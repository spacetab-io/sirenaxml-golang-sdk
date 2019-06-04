package structs

import "encoding/xml"

// CitiesResponse is a response to all cities request
type CitiesResponse struct {
	Answer  CitiesAnswer `xml:"answer"`
	XMLName xml.Name     `xml:"sirena" json:"-"`
}

// CitiesAnswer is an <answer> section in all cities response
type CitiesAnswer struct {
	Cities CitiesAnswerDetails `xml:"describe"`
}

// CitiesAnswerDetails is a <describe> section in all cities response
type CitiesAnswerDetails struct {
	Data  []CitiesAnswerData `xml:"data"`
	Error *Error             `xml:"error,omitempty"`
}

// CitiesAnswerData is a <data> section in all cities response
type CitiesAnswerData struct {
	Code     []CitiesAnswerDataCode    `xml:"code" json:"code"`
	Name     []CitiesAnswerDataName    `xml:"name" json:"name"`
	Country  []CitiesAnswerDataCountry `xml:"country" json:"country"`
	Timezone *CitiesAnswerDataTimezone `xml:"timezone" json:"timezone"`
}

// CitiesAnswerDataCode represents <code> entry in <data> section
type CitiesAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// CitiesAnswerDataName represents <name> entry in <data> section
type CitiesAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// CitiesAnswerDataCountry represents <country> entry in <data> section
type CitiesAnswerDataCountry struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// CitiesAnswerDataTimezone represents <timezone> entry in <data> section
type CitiesAnswerDataTimezone struct {
	ID    string `xml:"id,attr" json:"id"`
	Value string `xml:",chardata" json:"value"`
}
