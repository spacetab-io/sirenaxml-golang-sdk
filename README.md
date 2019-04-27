sirena-xml-sdk
--------------

Sirena XML connector written on golang

## Usage

```go
package main

import (
	"log"
	"os"
	"strconv"
	
	"github.com/microparts/logs-go"
	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/sdk"
	"github.com/tmconsulting/sirenaxml-golang-sdk/service"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func main() {
	clientID, _ := string2Uint16(os.Getenv("CLIENT_ID"))
	sc := &configuration.SirenaConfig{
		ClientID:                 clientID,
		Host:                     os.Getenv("HOST"),
		Port:                     os.Getenv("PORT"),
		ClientPublicKeyFile:      os.Getenv("CLIENT_PUBLIC_KEY"),
		ClientPrivateKeyFile:     os.Getenv("CLIENT_PRIVATE_KEY"),
		ServerPublicKeyFile:      os.Getenv("SERVER_PUBLIC_KEY"),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		KeysPath:                 os.Getenv("KEYS_PATH"),
		UseSymmetricKeyCrypt:     true,
		ZipRequests:              true,
		ZipResponses:             true,
	}
	lc := &logs.Config{
		Level:  "info",
		Format: "text",
	}
  	
	sdkClient, err := sdk.NewClient(sc, lc)
	if err != nil {
		log.Fatal(err)
	}
	srv := service.NewSKD(sdkClient)
	
	availabiliteReq := &structs.AvailabilityRequest{
		Query: structs.AvailabilityRequestQuery{
			Availability: structs.Availability{
				Departure: "MOW",
				Arrival:   "LED",
				AnswerParams: structs.AvailabilityAnswerParams{
					ShowFlighttime: true,
				},
			},
		},
	}
	response, err := srv.Avalability(availabiliteReq)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Print(response)
}


// String2Uint16 converts string to uint16
func string2Uint16(s string) (uint16, error) {
	b, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(b), nil
}
```

## Tests

Fill `.env` file from `test.env-example` and run tests:

	make test

## Licence

This software is provided under [MIT License](LICENSE).