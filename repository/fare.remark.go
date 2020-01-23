package repository

import (
	"encoding/xml"

	sirena "github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func (r Repository) GetFareRemark(logAttributes map[string]string, company, uptcode string, upt sirena.Upt) (*sirena.FareRemarkResponse, error) {
	params := sirena.FareRemarkRequestParams{
		Cat: 0,
		Upt: upt,
	}

	answerParams := sirena.FareRemarkAnswerParams{
		Lang: "RU",
	}
	query := sirena.FareRemarkRequestQuery{
		FareRemark: sirena.FareRemark{
			Company:       company,
			Code:          uptcode,
			RequestParams: params,
			AnswerParams:  answerParams,
		},
	}
	request := sirena.FareRemarkRequest{
		Query: query,
	}

	requestBytes, err := xml.MarshalIndent(&request, "  ", "    ")
	if err != nil {
		return nil, err
	}

	sirenaFareRemarkResponseXML, err := r.Transport.Request(requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	var sirenaFareRemarkResponse sirena.FareRemarkResponse
	if err = xml.Unmarshal(sirenaFareRemarkResponseXML, &sirenaFareRemarkResponse); err != nil {
		return nil, err
	}

	return &sirenaFareRemarkResponse, nil
}
