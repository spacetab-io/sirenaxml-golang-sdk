package sirena

import "encoding/xml"

// OrderRequest is a <order> request
type OrderRequest struct {
	Query   OrderRequestQuery `xml:"query"`
	XMLName xml.Name          `xml:"sirena"`
}

// OrderRequestQuery is a <query> section in <order> request
type OrderRequestQuery struct {
	Order Order `xml:"order"`
}

// Order is a body of <order> request
type Order struct {
	Regnum        string             `xml:"regnum"`
	Surname       string             `xml:"surname"`
	RequestParams OrderRequestParams `xml:"request_params"`
	AnswerParams  OrderAnswerParams  `xml:"answer_params"`
}

// OrderRequestParams is a <request_params> section in <order> request
type OrderRequestParams struct {
	TickSer           string `xml:"tick_ser"`
	NoPricing         bool   `xml:"no_pricing"`
	PrevPricingParams bool   `xml:"prev_pricing_params"`
	Formpay           string `xml:"formpay"`
}

// OrderAnswerParams is a <answer_params> section in <order> request
type OrderAnswerParams struct {
	Tickinfo           bool `xml:"tickinfo"`
	ShowTickinfoAgency bool `xml:"show_tickinfo_agency"`
	ShowActions        bool `xml:"show_actions"`
	AddCommonStatus    bool `xml:"add_common_status"`
	ShowUptRec         bool `xml:"show_upt_rec"`
	AddRemarks         bool `xml:"add_remarks"`
	AddSsr             bool `xml:"add_ssr"`
	AddPaycode         bool `xml:"add_paycode"`
	ShowErsp           bool `xml:"show_ersp"`
	ShowInsuranceInfo  bool `xml:"show_insurance_info"`
	ShowZh             bool `xml:"show_zh"`
	AddRemoteRecloc    bool `xml:"add_remote_recloc"`
	ShowComission      bool `xml:"show_comission"`
	ShowBagNorm        bool `xml:"show_bag_norm"`
}
