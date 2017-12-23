package sirena

import "encoding/xml"

type VehiclesResponse struct {
	Answer  VehiclesAnswer `xml:"answer"`
	XMLName xml.Name       `xml:"sirena" json:"-"`
}

type VehiclesAnswer struct {
	Vehicles VehiclesAnswerDetails `xml:"describe"`
}

type VehiclesAnswerDetails struct {
	Data []VehiclesAnswerData `xml:"data"`
}

type VehiclesAnswerData struct {
	Code []VehiclesAnswerDataCode `xml:"code"`
	Name []VehiclesAnswerDataName `xml:"name"`
}

type VehiclesAnswerDataCode struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type VehiclesAnswerDataName struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}
