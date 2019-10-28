package structs

import "encoding/xml"

type ViewFlownStatusRequest struct {
	Query   ViewFlownStatusQuery `xml:"query"`
	XMLName xml.Name             `xml:"sirena"`
}

type ViewFlownStatus struct {
	Regnum  string `xml:"regnum"`
	Surname string `xml:"surname"`
}

type ViewFlownStatusQuery struct {
	ViewFlownStatus ViewFlownStatus `xml:"view_flown_status"`
}
