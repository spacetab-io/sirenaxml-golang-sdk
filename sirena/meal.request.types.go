package sirena

import "encoding/xml"

// MealsRequest is a Sirena <describe> request for all meals
type MealsRequest struct {
	Query   MealsRequestQuery `xml:"query"`
	XMLName xml.Name          `xml:"sirena"`
}

// MealsRequestQuery is a <query> section in  all meals request
type MealsRequestQuery struct {
	Meals Meals `xml:"describe"`
}

// Meals is a <describe> section in all meals request
type Meals struct {
	Data string `xml:"data"`
	Code string `xml:"code,omitempty"`
}
