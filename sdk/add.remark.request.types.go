package sdk

import "encoding/xml"

// AddRemarkRequest is a <add_remark> request
type AddRemarkRequest struct {
	Query   AddRemarkRequestQuery `xml:"query"`
	XMLName xml.Name              `xml:"sirena"`
}

// AddRemarkRequestQuery is a <query> section in <add_remark> request
type AddRemarkRequestQuery struct {
	AddRemark AddRemark `xml:"add_remark"`
}

// AddRemark is a body of <add_remark> request
type AddRemark struct {
	Regnum  string `xml:"regnum"`
	Surname string `xml:"surname"`
	Type    string `xml:"type"`
	Remark  string `xml:"remark"`
}
