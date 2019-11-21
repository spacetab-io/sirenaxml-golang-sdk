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
	ID               int                   `xml:"id,attr"`
	Company          string                `xml:"company"`
	Num              string                `xml:"num"`
	OperatingCompany string                `xml:"operating_company"`
	Flight           string                `xml:"flight"`
	Origin           PricingAnswerLocation `xml:"origin"`
	Destination      PricingAnswerLocation `xml:"destination"`
	DeptDate         string                `xml:"deptdate"`
	ArrvDate         string                `xml:"arrvdate"`
	DeptTime         string                `xml:"depttime"`
	ArrvTime         string                `xml:"arrvtime"`
	Airplane         string                `xml:"airplane"`
	FlightTime       string                `xml:"flightTime"`
}

type PricingAnswerVariant struct {
	FlightGroups []PricingAnswerVariantFlightGroup `xml:"flights"`
	Directions   []PricingAnswerVariantDirection   `xml:"direction"`
	Total        PricingAnswerVariantTotal         `xml:"variant_total"`
}

// GetPaxBaseCost func return base cost of variant for transferred passenger type
func (p *PricingAnswerVariant) GetPaxBaseCost(paxType string) *float64 {

	// Passenger total cost will be add from all direction objects
	var paxBaseCost = new(float64)

	// Variant object contains objects of all passengers.
	// Since it is necessary to add the price for only one passenger of the transferred type, therefore passengerID is determined and the price is added only for a passenger of that passengerID.
	var passengerID int

	for _, price := range p.Directions[0].Prices {
		if price.Code == paxType {

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
			if price.Code == paxType && price.PassengerID == passengerID {


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
		if price.Code == paxType {

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
			if price.Code == paxType && price.PassengerID == passengerID {

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
		if price.Code == paxType {

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
			if price.Code == paxType && price.PassengerID == passengerID {

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
		if price.Code == paxType {

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
			if price.Code == paxType && price.PassengerID == passengerID {

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
		if price.Code == paxType {

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
			if price.Code == paxType && price.PassengerID == passengerID {

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
	Flight []PricingAnswerVariantFlight `xml:"flight"`
}

type PricingAnswerVariantFlight struct {
	ID         int    `xml:"id,attr"`
	Num        int    `xml:"num,attr"`
	SubClass   string `xml:"subclass,attr"`
	BaseClass  string `xml:"baseclass,attr"`
	Available  int    `xml:"available,attr"`
	SegmentNum string `xml:"iSegmentNum,attr"`
	Cabin      string `xml:"cabin,attr"`
}

type PricingAnswerVariantDirection struct {
	Num    int                   `xml:"num,attr"`
	Prices []*PricingAnswerPrice `xml:"price"`
}

type PricingAnswerLocation struct {
	City     string `xml:"city,attr"`
	Terminal string `xml:"terminal,attr"`
	Value    string `xml:",chardata"`
}

type PricingAnswerPrice struct {
	Upt               PriceUpt                `xml:"upt"`
	Fare              *PricingAnswerPriceFare `xml:"fare"`
	Taxes             []PricingAnswerPriceTax `xml:"tax"`
	Vat               *Vat                    `xml:"vat"`
	Baggage           string                  `xml:"baggage,attr"`
	ValidatingCompany string                  `xml:"validating_company,attr"`
	OrigCode          string                  `xml:"orig_code,attr"`
	Brand             string                  `xml:"brand,attr"`
	Total             float64                 `xml:"total"`
	Currency          string                  `xml:"currency,attr"`
	PassengerID       int                     `xml:"passenger-id,attr"`
	Code              string                  `xml:"code,attr"`
	// Count             int                     `xml:"count,attr"`
	// Ticket            string                  `xml:"ticket,attr"`
	// FC                string                  `xml:"fc,attr"`
	// DocID             string                  `xml:"doc_id,attr"`
	// ACCode            string                  `xml:"accode,attr"`
	// FOP               string                  `xml:"fop,attr"`
	// OrigID            int                     `xml:"orig_id,attr"`
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
