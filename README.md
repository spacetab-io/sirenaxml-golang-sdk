sirena-xml-sdk
--------------

[![CircleCI](https://circleci.com/gh/tmconsulting/sirenaxml-golang-sdk.svg?style=shield)](https://circleci.com/gh/tmconsulting/sirenaxml-golang-sdk) [![codecov](https://codecov.io/gh/tmconsulting/sirenaxml-golang-sdk/graph/badge.svg)](https://codecov.io/gh/tmconsulting/sirenaxml-golang-sdk)

Sirena XML connector written on golang

## Usage

```go
package main

import (
	"log"
	"os"
	"strconv"
	
	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/sdk"
	"github.com/tmconsulting/sirenaxml-golang-sdk/service"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func main() {
	clientID, _ := string2Uint16(os.Getenv("CLIENT_ID"))
	sc := &sirena.Config{
		ClientID:                 clientID,
		Ip:                     os.Getenv("IP"),
		Port:                     os.Getenv("PORT"),
		ClientPublicKey:      []byte(os.Getenv("CLIENT_PUBLIC_KEY")),
		ClientPrivateKey:     []byte(os.Getenv("CLIENT_PRIVATE_KEY")),
		ServerPublicKey:      []byte(os.Getenv("SERVER_PUBLIC_KEY")),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		KeysPath:                 os.Getenv("KEYS_PATH"),
		ZippedMessaging:             true,
	}
  
	logger := logs.NewNullLog()
	sdkClient, err := sdk.NewClient(sc, logger)
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
```

## Tests

Pass config data in ENV and run tests:

	make test

## Licence

This software is provided under [MIT License](LICENSE).