package sirena

import "encoding/xml"

// VehiclesRequest is a Sirena <describe> request for all vehicles
type VehiclesRequest struct {
	Query   VehiclesRequestQuery `xml:"query"`
	XMLName xml.Name             `xml:"sirena"`
}

// VehiclesRequestQuery is a <query> section in all vehicles request
type VehiclesRequestQuery struct {
	Vehicles Vehicles `xml:"describe"`
}

// Vehicles is a <describe> section in all vehicles request
type Vehicles struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
