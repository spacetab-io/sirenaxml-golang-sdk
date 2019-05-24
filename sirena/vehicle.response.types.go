package sirena

import "encoding/xml"

// VehiclesResponse is a response to all vehicles request
type VehiclesResponse struct {
	Answer  VehiclesAnswer `xml:"answer"`
	XMLName xml.Name       `xml:"sirena" json:"-"`
}

// VehiclesAnswer is an <answer> section in all vehicles response
type VehiclesAnswer struct {
	Vehicles VehiclesAnswerDetails `xml:"describe"`
}

// VehiclesAnswerDetails is a <describe> section in all vehicles response
type VehiclesAnswerDetails struct {
	Data  []VehiclesAnswerData `xml:"data"`
	Error *Error               `xml:"error"`
}

// VehiclesAnswerData is a <data> section in all vehicles response
type VehiclesAnswerData struct {
	Code []VehiclesAnswerDataCode `xml:"code" json:"code"`
	Name []VehiclesAnswerDataName `xml:"name" json:"name"`
}

// VehiclesAnswerDataCode represents <code> entry in <data> section
type VehiclesAnswerDataCode struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}

// VehiclesAnswerDataName represents <name> entry in <data> section
type VehiclesAnswerDataName struct {
	Lang  string `xml:"lang,attr" json:"lang"`
	Value string `xml:",chardata" json:"value"`
}
