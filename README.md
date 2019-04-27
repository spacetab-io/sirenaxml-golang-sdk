sirena-xml-sdk
--------------

Sirena XML connector written on golang

## Usage

```go
package main

import (
	"log"
	"os"
	
	"github.com/microparts/logs-go"
	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/sirena"
	"github.com/tmconsulting/sirenaxml-golang-sdk/utils"
)

func main() {
	clientID, _ := utils.String2Uint16(os.Getenv("CLIENT_ID"))
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
  	
	client, err := sirena.NewClient(sc, lc)
	if err != nil {
		log.Fatal(err)
	}
	
	var keyInfoXML = []byte(`<?xml version="1.0" encoding="UTF-8"?><sirena><query><key_info/></query></sirena>`)

	request := &sirena.Request{
  		Message: keyInfoXML,
  		Header: sirena.NewHeader(&sirena.NewHeaderParams{
  			ClientID: client.Config.ClientID,
        MessageLength:  uint32(len(keyInfoXML)),
      }),
  	}

  response, err := client.Send(request)
  if err != nil {
    log.Fatal(err)
  }
  	
  log.Print(response)
}
```

## Tests

Fill `.env` file from `test.env-example` and run tests:

	make test

## Licence

This software is provided under [MIT License](LICENSE).