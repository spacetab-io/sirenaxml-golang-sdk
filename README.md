sirena-xml-sdk
--------------

[![CircleCI](https://circleci.com/gh/tmconsulting/sirenaxml-golang-sdk.svg?style=shield)](https://circleci.com/gh/tmconsulting/sirenaxml-golang-sdk) [![codecov](https://codecov.io/gh/tmconsulting/sirenaxml-golang-sdk/graph/badge.svg)](https://codecov.io/gh/tmconsulting/sirenaxml-golang-sdk)

Sirena XML connector written on golang

## Usage

1. Init new connection
2. connect
    - start listener
        - receive frame
        - check type
        - transform frame to message
        - save message
    - start sender
        - make message
        - set type
        - transform to frame
        - send frame
    - start sign key

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
	clientID, err := strconv.ParseUint(os.Getenv("CLIENT_ID"), 10, 16)
	if err != nil {
	  panic(err)
	}
	sc := &sirenaXML.Config{
		ClientID:                 uint16(clientID),
		Ip:                       os.Getenv("IP"),
		Environment:              os.Getenv("ENV"),
		ClientPublicKey:      	  os.Getenv("CLIENT_PUBLIC_KEY"),
		ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
		ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		ZippedMessaging:          true,
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

Feel free to help us to rise coverage!

## Contribution

Contribution, in any kind of way, is highly welcome!
It doesn't matter if you are not able to write code.
Creating issues or holding talks and help other people to use 
[sirenaxml-golang-sdk](https://github.com/tmconsulting/sirenaxml-golang-sdk) is contribution, too!

A few examples:

* Correct typos in the README / documentation
* Reporting bugs
* Implement a new feature or service
* Sharing the love if like to use [sirenaxml-golang-sdk](https://github.com/tmconsulting/sirenaxml-golang-sdk) and help people 
to get use to it

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

## License

SDK is released under the [MIT License](./LICENSE).