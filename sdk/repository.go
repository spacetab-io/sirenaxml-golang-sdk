package sdk

import (
	"encoding/xml"
	"net"
	"time"

	log "github.com/microparts/logs-go"
	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

// SirenaClient is a sirena client
type SirenaClient struct {
	net.Conn
	Key    []byte
	KeyID  uint32
	Config *configuration.SirenaConfig
}

// NewClient connects to Sirena (if not yet) and returns sirena client singleton
func NewClient(sc *configuration.SirenaConfig, lc *log.Config) (*SirenaClient, error) {
	err := sc.GetCerts()
	if err != nil {
		return nil, err
	}
	err = logs.Init(lc)
	if err != nil {
		return nil, errors.Wrap(err, "sirena client loggin init error")
	}
	conn, err := net.Dial("tcp", sc.GetSirenaAddr())
	if err != nil {
		return nil, errors.Wrap(err, "dial sirena addr error")
	}
	client := &SirenaClient{
		Conn:   conn,
		Key:    nil,
		Config: sc,
	}
	if client.Key == nil {
		// Create symmetric key
		if err := client.CreateAndSignKey(); err != nil {
			return nil, errors.Wrap(err, "creating and signing key error")
		}
		// Update key every 1 hour
		// @TODO что-то это нихера не очевидная ебулдень. Оно точно будет работать и будет работать корректно?
		go func() {
			for range time.Tick(time.Hour) {
				if err := client.CreateAndSignKey(); err != nil {
					logs.Log.Fatal("key updating error")
				}
			}
		}()
	}
	return client, nil
}

func (client *SirenaClient) GetAvailability(req []byte) (*structs.AvailabilityResponse, error) {
	// Create Sirena request
	request, err := client.NewRequest(req)
	if err != nil {
		return nil, err
	}
	// Send request to Sirena
	response, err := client.Send(request)
	if err != nil {
		return nil, err
	}

	var availabilityResponse structs.AvailabilityResponse
	if err := xml.Unmarshal(response.Message, &availabilityResponse); err != nil {
		return nil, err
	}

	return &availabilityResponse, nil
}
