package repository

import (
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/tmconsulting/sirenaxml-golang-sdk/publisher"
	"github.com/tmconsulting/sirenaxml-golang-sdk/service"
)

type Config struct {
	ProxyURL string
}

//Storage stow storage container
type Repository struct {
	p             *http.Client
	LogsPublisher publisher.Publisher
	Config        Config
	Transport     service.Storage
}

func NewStorage(transport service.Storage) (*Repository, error) {
	r := new(Repository)

	r.p = &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}

	r.Transport = transport

	err := r.ping()

	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrap(err, "ping proxy error")
	}

	return r, nil
}

// ping makes sure Sirena proxy is alive
func (r *Repository) ping() error {
	// testRequestBytes := []byte(`<?xml version="1.0" encoding="UTF-8"?><sirena><query><key_info/></query></sirena>`)
	// response, err := r.ProxySendRequest(testRequestBytes, "key_info", "key_info")
	// if err != nil {
	// 	return err
	// }
	// logs.Log.Debugf("%r", response)
	return nil
}
