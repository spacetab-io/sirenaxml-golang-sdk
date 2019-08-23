package errors

import (
	"encoding/xml"

	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func CheckErrors(responseError []byte, err error) (bool, *structs.Error, error) {
	if err != nil {
		return true, nil, err
	}
	if responseError != nil {
		var errAnswer structs.ErrorResponse
		err := xml.Unmarshal(responseError, &errAnswer)
		if err != nil {
			return true, nil, err
		}

		return true, &errAnswer.Answer.Error, nil
	}

	return false, nil, nil
}
