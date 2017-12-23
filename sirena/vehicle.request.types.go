package sirena

import "encoding/xml"

// VehiclesRequest is a <describe> request
type VehiclesRequest struct {
	Query   VehiclesRequestQuery `xml:"query"`
	XMLName xml.Name             `xml:"sirena"`
}

// VehiclesRequestQuery is a <query> section in <describe> request
type VehiclesRequestQuery struct {
	Vehicles Vehicles `xml:"describe"`
}

type Vehicles struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
