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
	"github.com/tmconsulting/sirena-config"
	"github.com/tmconsulting/sirenaxml-golang-sdk/sirena"
)

const keyInfoXML = `<?xml version="1.0" encoding="UTF-8"?>
<sirena>
  <query>
    <key_info/>
  </query>
</sirena>`

func main() {
	sc := &sirenaConfig.SirenaConfig{
  		ClientID:                 os.Getenv("CLIENT_ID"),
  		Host:                     os.Getenv("HOST"),
  		Port:                     os.Getenv("PORT"),
  		ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
  		ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
  		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
  		ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
  		KeysPath:                 os.Getenv("KEYS_PATH"),
  	}
  	lc := &logs.Config{
  		Level:  "info",
  		Format: "text",
  	}
  	
	client := sirena.NewClient(sc, lc, sirena.NewClientOptions{Test: true})
	request := &sirena.Request{
  		Message: []byte(keyInfoXML),
  	}
	var err error
  request.Header, err = sirena.NewHeader(sc, sirena.NewHeaderParams{
    Message: request.Message,
  })
  if err != nil {
    log.Fatal(err)
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

	go test ./... -v

## Licence

This software is provided under [MIT Licence](LICENCE).