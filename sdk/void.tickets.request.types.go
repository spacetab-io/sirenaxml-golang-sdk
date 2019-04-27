package sdk

import "encoding/xml"

// VoidTicketsRequest is a <void_tickets> request
type VoidTicketsRequest struct {
	Query   VoidTicketsRequestQuery `xml:"query"`
	XMLName xml.Name                `xml:"sirena"`
}

// VoidTicketsRequestQuery is a <query> section in <void_tickets> request
type VoidTicketsRequestQuery struct {
	VoidTickets VoidTickets `xml:"void_tickets"`
}

// VoidTickets is a body of <void_tickets> request
type VoidTickets struct {
	Regnum        string                   `xml:"regnum"`
	Surname       string                   `xml:"surname"`
	RequestParams VoidTicketsRequestParams `xml:"request_params"`
}

// VoidTicketsRequestParams is a <request_params> section in <order> request
type VoidTicketsRequestParams struct {
	ReturnSeats bool `xml:"return_seats"`
}
