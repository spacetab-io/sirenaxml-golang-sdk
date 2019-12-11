package structs

import (
	"encoding/xml"
)

type PricingResponse struct {
	Answer PricingAnswer `xml:"answer"`
}

type PricingAnswer struct {
	Pricing          *PricingAnswerPricing `xml:"pricing,omitempty"`
	PricingMonobrand *PricingAnswerPricing `xml:"pricing_mono_brand,omitempty"`
}

type PricingAnswerPricing struct {
	Variants []PricingAnswerVariant `xml:"variant"`
	Flights  []*PricingAnswerFlight `xml:"flight"`
	Error    *Error                 `xml:"error,omitempty"`
}

type PricingAnswerFlight struct {
	ID               string                `xml:"id,attr"`
	MarketingAirline string                `xml:"company"`
	FlightNumber     string                `xml:"num"`
	OperatingAirline string                `xml:"operating_company"`
	Flight           string                `xml:"flight"`
	Origin           PricingAnswerLocation `xml:"origin"`
	Destination      PricingAnswerLocation `xml:"destination"`
	DeptDate         string                `xml:"deptdate"`
	Legs             []FlightLeg           `xml:"legs>leg"`
	ArrvDate         string                `xml:"arrvdate"`
	DeptTime         string                `xml:"depttime"`
	ArrvTime         string                `xml:"arrvtime"`
	Airplane         string                `xml:"airplane"`
	FlightTime       string                `xml:"flightTime"`
	UPT18CatText     string                `xml:"cat_18_text,omitempty"`
}

// GetVariantFlightInfo func return flight info from variant flight
func (p *PricingAnswerFlight) GetVariantFlightInfo(variants []PricingAnswerVariant) *PricingAnswerVariantFlight {
	for _, variant := range variants {
		for _, flightGroup := range variant.FlightGroups {
			for _, flight := range flightGroup.Flight {
				if p.ID == flight.ID {
					return &flight
				}
			}
		}
	}

	return nil
}

type PricingAnswerVariant struct {
	FlightGroups []PricingAnswerVariantFlightGroup `xml:"flights"`
	Directions   []PricingAnswerVariantDirection   `xml:"direction"`
	Total        PricingAnswerVariantTotal         `xml:"variant_total"`
}

func (p *PricingAnswerVariant) GetVariantDirections() []PricingAnswerVariantDirection {
	return p.Directions
}

// GetPaxBaseCost func return base cost of variant for transferred passenger type
func (p *PricingAnswerVariant) GetPaxBaseCost(paxType string) *float64 {

	// Passenger total cost will be add from all direction objects
	var paxBaseCost = new(float64)

	// Variant object contains objects of all passengers.
	// Since it is necessary to add the price for only one passenger of the transferred type, therefore passengerID is determined and the price is added only for a passenger of that passengerID.
	var passengerID int

	for _, price := range p.Directions[0].Prices {
		if price.PaxType == paxType {

			// Attach ID of first passenger of appropriate type to passengerID
			passengerID = price.PassengerID
			break
		}
	}

	for _, direction := range p.Directions {

		// Prices object contains cost info related to type of pax
		// Add passenger price form all direction objects
		for _, price := range direction.Prices {

			// Check if it is needed passenger type
			if price.PaxType == paxType && price.PassengerID == passengerID {

				*paxBaseCost += price.Fare.Total
			}
		}
	}

	return paxBaseCost
}

// GetVariantBaseCost func return base cost of variant for all variant
func (p *PricingAnswerVariant) GetVariantBaseCost() *float64 {

	// Passenger total cost will be add from all direction objects
	var variantBaseCost = new(float64)

	for _, direction := range p.Directions {

		// Prices object contains cost info related to type of pax
		// Add passenger price form all direction objects
		for _, price := range direction.Prices {

			*variantBaseCost += price.Fare.Total
		}
	}

	return variantBaseCost
}

// GetPaxTotalCost func return total cost of variant for transferred passenger type
func (p *PricingAnswerVariant) GetPaxTotalCost(paxType string) *float64 {

	// Passenger total cost will be add from all direction objects
	var paxTotalCost = new(float64)

	// Variant object contains objects of all passengers.
	// Since it is necessary to add the price for only one passenger of the transferred type, therefore passengerID is determined and the price is added only for a passenger of that passengerID.
	var passengerID int

	for _, price := range p.Directions[0].Prices {
		if price.PaxType == paxType {

			// Attach ID of first passenger of appropriate type to passengerID
			passengerID = price.PassengerID
			break
		}
	}

	for _, direction := range p.Directions {

		// Prices object contains cost info related to type of pax
		// Add passenger price form all direction objects
		for _, price := range direction.Prices {

			// Check if it is needed passenger type
			if price.PaxType == paxType && price.PassengerID == passengerID {

				//allocate a new zero-valued paxTotalCost

				*paxTotalCost += price.Total
			}
		}
	}

	return paxTotalCost
}

// GetPaxTaxesCost func return taxes cost of variant for transferred passenger type
func (p *PricingAnswerVariant) GetPaxTaxesCost(paxType string) *float64 {

	// Passenger total cost will be add from all direction objects
	var paxTaxesCost = new(float64)

	// Variant object contains objects of all passengers.
	// Since it is necessary to add the price for only one passenger of the transferred type, therefore passengerID is determined and the price is added only for a passenger of that passengerID.
	var passengerID int

	for _, price := range p.Directions[0].Prices {
		if price.PaxType == paxType {

			// Attach ID of first passenger of appropriate type to passengerID
			passengerID = price.PassengerID
			break
		}
	}

	for _, direction := range p.Directions {

		// Prices object contains cost info related to type of pax
		// Add passenger price form all direction objects
		for _, price := range direction.Prices {

			// Check if it is needed passenger type
			if price.PaxType == paxType && price.PassengerID == passengerID {

				//allocate a new zero-valued paxTotalCost
				//paxTaxesCost = new(float64)

				for _, tax := range price.Taxes {
					*paxTaxesCost += tax.Total
				}
			}
		}
	}

	return paxTaxesCost
}

// GetVariantTaxesCost func return taxes cost of variant for all variant
func (p *PricingAnswerVariant) GetVariantTaxesCost() *float64 {

	// Passenger total cost will be add from all direction objects
	var taxesCost = new(float64)

	for _, direction := range p.Directions {

		// Prices object contains cost info related to type of pax
		// Add passenger price form all direction objects
		for _, price := range direction.Prices {

			for _, tax := range price.Taxes {
				*taxesCost += tax.Total
			}
		}
	}

	return taxesCost
}

// GetPaxTaxesRow func return row taxes slice of variant for transferred passenger type
func (p *PricingAnswerVariant) GetPaxTaxesRow(paxType string) []PricingAnswerPriceTax {

	// Passenger total cost will be add from all direction objects
	var paxTaxes []PricingAnswerPriceTax

	// Variant object contains objects of all passengers.
	// Since it is necessary to add the price for only one passenger of the transferred type, therefore passengerID is determined and the price is added only for a passenger of that passengerID.
	var passengerID int

	for _, price := range p.Directions[0].Prices {
		if price.PaxType == paxType {

			// Attach ID of first passenger of appropriate type to passengerID
			passengerID = price.PassengerID
			break
		}
	}

	for _, direction := range p.Directions {

		// Prices object contains cost info related to type of pax
		// Add passenger price form all direction objects
		for _, price := range direction.Prices {

			// Check if it is needed passenger type
			if price.PaxType == paxType && price.PassengerID == passengerID {

			TAXES_LOOP:
				for _, tax := range price.Taxes {

					for _, containsPax := range paxTaxes {

						if tax.Total == containsPax.Total {
							continue TAXES_LOOP
						}
					}

					paxTaxes = append(paxTaxes, tax)
				}
			}
		}
	}

	return paxTaxes
}

// GetPaxVatsRow func return row vats slice of variant for transferred passenger type
func (p *PricingAnswerVariant) GetPaxVatsRow(paxType string) []*Vat {

	// Passenger total cost will be add from all direction objects
	var paxVats []*Vat

	// Variant object contains objects of all passengers.
	// Since it is necessary to add the price for only one passenger of the transferred type, therefore passengerID is determined and the price is added only for a passenger of that passengerID.
	var passengerID int

	for _, price := range p.Directions[0].Prices {
		if price.PaxType == paxType {

			// Attach ID of first passenger of appropriate type to passengerID
			passengerID = price.PassengerID
			break
		}
	}

	for _, direction := range p.Directions {

		// Prices object contains cost info related to type of pax
		// Add passenger price form all direction objects
		for _, price := range direction.Prices {

			// Check if it is needed passenger type
			if price.PaxType == paxType && price.PassengerID == passengerID {

				if price.Vat != nil && price.Vat.Vats != nil {

					paxVats = append(paxVats, price.Vat)
				}
			}
		}
	}

	return paxVats
}

func (p *PricingAnswerVariant) GetVariantTotalCost() *PricingAnswerVariantTotal {

	return &p.Total
}

func (p *PricingAnswerVariant) GetVariantPriceCurrency() string {

	return p.Directions[0].Prices[0].Currency
}

type PricingAnswerVariantTotal struct {
	Currency string  `xml:"currency,attr"`
	Total    float64 `xml:",chardata"`
}

type PricingAnswerVariantFlightGroup struct {
	BookingTimeLimit      string                       `xml:"utc_timelimit,attr,omitempty"`
	NeedLatinRegistration bool                         `xml:"latin_registration,attr,omitempty"`
	ETPossible            bool                         `xml:"et_possible,attr,omitempty"`
	ETBlanks              bool                         `xml:"et_blanks,attr,omitempty"`
	Flight                []PricingAnswerVariantFlight `xml:"flight"`
}

func (p *PricingAnswerVariantFlightGroup) GetBrandChecked(variant PricingAnswerVariant, paxType string) bool {

	var flightHaveBrand bool

	for _, flight := range p.Flight {
		if flight.GetVariantPricing(variant, paxType).BrandCode != "" {
			flightHaveBrand = true
		}
	}

	return flightHaveBrand
}

func (f *PricingAnswerVariantFlight) GetFlightInfo(flights []*PricingAnswerFlight) *PricingAnswerFlight {
	for _, flight := range flights {

		//spew.Dump(" \n flight.ID: \n", flight.ID)
		//spew.Dump(" \n f.ID: \n", f.ID)
		if f.ID == flight.ID {
			return flight
		}
	}

	return nil
}

func findFlight(flights []PricingAnswerVariantFlight, f *PricingAnswerVariantFlight) bool {
	// Find variant that contain the current PricingAnswerVariantFlight
	for _, flight := range flights {
		if f.ID == flight.ID {

			// Once variant that contained a PricingAnswerVariantFlight had found, starting find PAX data attached to PricingAnswerVariantFlight
			return true

		}
	}

	return false
}

func (f *PricingAnswerVariantFlight) GetPaxInfoFromFlight(variants []PricingAnswerVariant, paxType string) *PricingAnswerPrice {
	for _, variant := range variants {
		for _, flights := range variant.FlightGroups {

			// Find variant that contain the current PricingAnswerVariantFlight
			if findFlight(flights.Flight, f) {

				for _, direction := range variant.Directions {
					for _, price := range direction.Prices {

						if price.PaxType == paxType {

							return price
						}
					}
				}
			}
		}
	}

	return nil
}

func (f *PricingAnswerVariantFlight) GetVariantPricing(variant PricingAnswerVariant, paxType string) *PricingAnswerPrice {

	for _, direction := range variant.Directions {

		if direction.Num == f.Num {

			for _, price := range direction.Prices {

				if price.PaxType == paxType {

					return price
				}
			}
		}
	}

	return nil
}

type PricingAnswerVariantFlight struct {
	ID         string `xml:"id,attr"`
	Num        int    `xml:"num,attr"`
	Class      string `xml:"class,attr"`
	SubClass   string `xml:"subclass,attr"`
	BaseClass  string `xml:"baseclass,attr"`
	Available  int    `xml:"available,attr"`
	SegmentNum string `xml:"iSegmentNum,attr"`
	Cabin      string `xml:"cabin,attr"`
}

type FlightLeg struct {
	Airplane string              `xml:"airplane,attr"`
	Dep      PNRSegmentLegDepArr `xml:"dep"`
	Arr      PNRSegmentLegDepArr `xml:"arr"`
}

// PNRSegmentLegDepArr is <sep> and <arr> entries in <leg> section
type FlightLegDepArr struct {
	TimeLocal string `xml:"time_local,attr"`
	TimeUTC   string `xml:"time_utc,attr"`
	Term      string `xml:"term,attr"`
	Value     string `xml:",chardata"`
}

type PricingAnswerVariantDirection struct {
	Num            int                   `xml:"num,attr"`
	RequestedBrand string                `xml:"requested_brand,attr,omitempty"`
	Prices         []*PricingAnswerPrice `xml:"price"`
}

func (p *PricingAnswerVariantDirection) GetDirectionsFlights(variant *PricingAnswerVariant) []PricingAnswerVariantFlight {

	var directionFlights []PricingAnswerVariantFlight

	for _, flightGroups := range variant.FlightGroups {
		for _, flight := range flightGroups.Flight {
			if p.Num == flight.Num {
				directionFlights = append(directionFlights, flight)
			}
		}
	}

	return nil
}

func (p *PricingAnswerVariant) GetDirectionsFlights(directionNum int) [][]PricingAnswerVariantFlight {
	// Declare slice for flightGroups of one direction
	var directionFlightGroups [][]PricingAnswerVariantFlight

FLIGHTGROUPS_LABEL:
	for _, flightGroups := range p.FlightGroups {

		var directionFlights []PricingAnswerVariantFlight

		for _, flg := range directionFlightGroups {
			for _, fl := range flg {
				for _, flight := range flightGroups.Flight {
					if fl.ID == flight.ID {

						continue FLIGHTGROUPS_LABEL
					}
				}
			}
		}

		for _, flight := range flightGroups.Flight {

			// Declare slice for flights of one direction
			//var directionFlights []PricingAnswerVariantFlight
			if directionNum == flight.Num {

				directionFlights = append(directionFlights, flight)
			}

			// check if directionFlights have any items
			if len(directionFlights) == 0 {

				continue
			}

		}

		directionFlightGroups = append(directionFlightGroups, directionFlights)

	}

	return directionFlightGroups
}

func (p *PricingAnswerVariantDirection) GetDirectionsFlightGroups(variant *PricingAnswerVariant) [][]PricingAnswerVariantFlight {

	// Declare slice for flightGroups of one direction
	var directionFlightGroups [][]PricingAnswerVariantFlight

FLIGHTGROUPS_LABEL:
	for _, flightGroups := range variant.FlightGroups {

		var directionFlights []PricingAnswerVariantFlight

		for _, flg := range directionFlightGroups {
			for _, fl := range flg {
				for _, flight := range flightGroups.Flight {
					if fl.ID == flight.ID {

						continue FLIGHTGROUPS_LABEL
					}
				}
			}
		}

		for _, flight := range flightGroups.Flight {

			// Declare slice for flights of one direction
			//var directionFlights []PricingAnswerVariantFlight
			if p.Num == flight.Num {

				directionFlights = append(directionFlights, flight)
			}

			// check if directionFlights have any items
			if len(directionFlights) == 0 {

				continue
			}

		}

		directionFlightGroups = append(directionFlightGroups, directionFlights)
	}

	return directionFlightGroups
}

func (p *PricingAnswerVariantDirection) GetDirectionsVariants(variant *PricingAnswerVariant) [][]PricingAnswerVariantFlight {

	// Declare slice for flightGroups of one direction
	var directionFlightGroups [][]PricingAnswerVariantFlight
	for _, flightGroups := range variant.FlightGroups {

		var directionFlights []PricingAnswerVariantFlight
		for _, flight := range flightGroups.Flight {

			// Declare slice for flights of one direction
			//var directionFlights []PricingAnswerVariantFlight

			if p.Num == flight.Num {
				directionFlights = append(directionFlights, flight)
			}

			// check if directionFlights have any items
			if len(directionFlights) == 0 {

				continue
			}

		}
		directionFlightGroups = append(directionFlightGroups, directionFlights)
	}

	return directionFlightGroups
}

type PricingAnswerLocation struct {
	City     string `xml:"city,attr"`
	Terminal string `xml:"terminal,attr"`
	Value    string `xml:",chardata"`
}

type PricingAnswerPrice struct {
	Brand             string                  `xml:"brand,attr"`
	Baggage           string                  `xml:"baggage,attr"`
	ValidatingAirline string                  `xml:"validating_company,attr"`
	OriginalPaxType   string                  `xml:"orig_code,attr"`
	//BrandCode         string                  `xml:"brand,attr"`
	Currency          string                  `xml:"currency,attr"`
	PassengerID       int                     `xml:"passenger-id,attr"`
	PaxType           string                  `xml:"code,attr"`
	IsRoundTrip       bool                    `xml:"rt,attr,omitempty"`
	FormPay           FormPay                 `xml:"fop,attr,omitempty"`
	Upt               PriceUpt                `xml:"upt"`
	Fare              *PricingAnswerPriceFare `xml:"fare"`
	Taxes             []PricingAnswerPriceTax `xml:"tax"`
	Vat               *Vat                    `xml:"vat"`
	Total             float64                 `xml:"total"`
	UPT18CatText      string                  `xml:"cat18_text,omitempty"`
	// Count             int                     `xml:"count,attr"`
	// Ticket            string                  `xml:"ticket,attr"`
	// FC                string                  `xml:"fc,attr"`
	// DocID             string                  `xml:"doc_id,attr"`
	// ACCode            string                  `xml:"accode,attr"`
	// FOP               string                  `xml:"fop,attr"`
	// OrigID            int                     `xml:"orig_id,attr"`
}

func (p *PricingAnswerPrice) GetFare() *PricingAnswerPriceFare {

	return p.Fare
}

func (v *Vat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v.Fare = start.Name.Local
	v.Vats = start.Attr
	return d.Skip()
}

type Vat struct {
	Fare string `xml:"fare,attr"`
	Vats []xml.Attr
}

type PricingAnswerPriceFare struct {
	Remark      string  `xml:"remark,attr"`
	FareExpDate string  `xml:"fare_expdate,attr"`
	Code        string  `xml:"code,attr"`
	BaseCode    string  `xml:"base_code,attr"`
	Total       float64 `xml:",chardata"`
}

type PricingAnswerPriceTax struct {
	Code  string  `xml:"code,attr"`
	Owner string  `xml:"owner,attr"`
	Total float64 `xml:",chardata"`
}

type PriceUpt struct {
	Idar1     string `xml:"idar1"`
	AddonIda  string `xml:"addon_ida" json:"addon_ida"`
	Ntrip     string `xml:"ntrip"`
	Nvr       string `xml:"nvr"`
	Ftnt      string `xml:"ftnt"`
	CodeUpt   string `xml:"code_upt" json:"code_upt"`
	Tariff    string `xml:"tariff"`
	MainAwk   string `xml:"main_awk" json:"main_awk"`
	Cat       string `xml:"cat"`
	Vcat      string `xml:"vcat"`
	City1     string `xml:"city1"`
	City2     string `xml:"city2"`
	Dport     string `xml:"dport"`
	Aport     string `xml:"aport"`
	BaseFare  string `xml:"base_fare" json:"base_fare"`
	Iit       string `xml:"iit"`
	Owrt      string `xml:"owrt"`
	Ddate     string `xml:"ddate"`
	Fdate     string `xml:"fdate"`
	DelivType string `xml:"deliv_type" json:"deliv_type"`
	F0        string `xml:"f0"`
	F1        string `xml:"f1"`
	F2        string `xml:"f2"`
	F3        string `xml:"f3"`
	FlAwk     string `xml:"fl_awk" json:"fl_awk"`
}
