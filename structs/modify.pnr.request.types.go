package structs

import (
	"encoding/xml"
)

// ModifyPNRRequest is a <modify_pnr> request
type ModifyPNRRequest struct {
	Query   ModifyPNRQuery `xml:"query"`
	XMLName xml.Name       `xml:"sirena"`
}

// ModifyPNRQuery is a <query> section in <modify_pnr> request
type ModifyPNRQuery struct {
	ModifyPNR ModifyPNR `xml:"modify_pnr"`
}

// ModifyPNR is a body of <modify_pnr> request
type ModifyPNR struct {
	Version      int                   `xml:"version"`
	Regnum       string                `xml:"regnum"`
	Surname      string                `xml:"surname"`
	AddParams    ModifyPNRAddParams    `xml:"add,omitempty"`
	RemoveParams ModifyPNRRemoveParams `xml:"remove"`
}

// ModifyPNRAddParams is a <add> section in <modify_pnr> request
type ModifyPNRAddParams struct {
	Contact      []ModifyPNRContact      `xml:"contact,omitempty"`
	PassDocument []ModifyPNRPassDocument `xml:"pass_document,omitempty"`
}

type ModifyPNRRemoveParams struct {
	Ssr []Ssr `xml:"ssr"`
}

// ModifyPNRContact is <contact> element for <add> section
type ModifyPNRContact struct {
	Surname string `xml:"surname,attr,omitempty"`
	Name    string `xml:"name,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Comment string `xml:"comment,attr,omitempty"`
	Value   string `xml:",chardata"`
}

// ModifyPNRPassDocument is <pass_document> element for <add> section
type ModifyPNRPassDocument struct {
	Surname    string `xml:"surname,attr"`
	Name       string `xml:"name,attr"`
	Doccode    string `xml:"doccode,attr"`
	DocCountry string `xml:"doc_country,attr"`
	PspExpire  string `xml:"pspexpire,attr,omitempty"`
	Value      string `xml:",chardata"`
}
