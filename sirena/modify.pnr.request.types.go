package sirena

import "encoding/xml"

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
	Regnum    ModifyPNRRegnum    `xml:"regnum"`
	Surname   string             `xml:"surname"`
	Modify    ModifyPNRModify    `xml:"modify,omitempty"`
	AddParams ModifyPNRAddParams `xml:"add,omitempty"`
}

// ModifyPNRRegnum is a <regnum> element of <modify_pnr> request
type ModifyPNRRegnum struct {
	Version string `xml:"version,attr"`
	Value   string `xml:",chardata"`
}

// ModifyPNRRegnum is a <regnum> element of <modify_pnr> request
type ModifyPNRModify struct {
	Passenger ModifyPNRModifyPassenger `xml:"passenger,omitempty"`
	Contacts  BookingAnswerContacts    `xml:"contacts,omitempty"`
}

type ModifyPNRModifyPassenger struct {
	PassID string `xml:"pass_id,attr,omitempty"`
	// LeadPass    bool                 `xml:"lead_pass,attr,omitempty"`
	Name      string `xml:"name,omitempty"`
	Surname   string `xml:"surname"`
	Category  string `xml:"category",omitempty`
	Sex       string `xml:"sex",omitempty`
	Birthdate string `xml:"birthdate",omitempty`
	// Age         int                  `xml:"age",omitempty`
	DocCode     string `xml:"doccode",omitempty`
	Doc         string `xml:"doc",omitempty`
	PspExpire   string `xml:"pspexpire",omitempty`
	DocCountry  string `xml:"doc_country",omitempty`
	Nationality string `xml:"nationality",omitempty`
	Residence   string `xml:"residence",omitempty`
	// Contacts    []Contact            `xml:"contacts>contact",omitempty`
}

// ModifyPNRAddParams is a <add> section in <modify_pnr> request
type ModifyPNRAddParams struct {
	Contact      []ModifyPNRContact        `xml:"contact,omitempty"`
	PassDocument []ModifyPNRPassDocument   `xml:"pass_document,omitempty"`
	Passenger    []BookingRequestPassenger `xml:"passenger,omitempty"`
}

//ModifyPNRContact is <contact> element for <add> section
type ModifyPNRContact struct {
	Surname string `xml:"surname,attr,omitempty"`
	Name    string `xml:"name,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Comment string `xml:"comment,attr,omitempty"`
	Value   string `xml:",chardata"`
}

//ModifyPNRPassDocument is <pass_document> element for <add> section
type ModifyPNRPassDocument struct {
	Surname    string `xml:"surname,attr"`
	Name       string `xml:"name,attr"`
	Doccode    string `xml:"doccode,attr"`
	DocCountry string `xml:"doc_country,attr"`
	PspExpire  string `xml:"pspexpire,attr,omitempty"`
	Value      string `xml:",chardata"`
}
