package repository

import (
	"bytes"
	"github.com/pkg/errors"
	sirenaXML "github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/publisher"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//Storage stow storage container
type Repository struct {
	p             *http.Client
	LogsPublisher publisher.Publisher
	Config        sirenaXML.Config
}

func NewStorage(logsPublisher publisher.Publisher) (*Repository, error) {
	s := new(Repository)

	s.p = &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}

	s.LogsPublisher = logsPublisher

	config, err := sirenaXML.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	s.Config = *config

	err = s.ping()

	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrap(err, "ping proxy error")
	}

	return s, nil
}

// ping makes sure Sirena proxy is alive
func (s *Repository) ping() error {
	// testRequestBytes := []byte(`<?xml version="1.0" encoding="UTF-8"?><sirena><query><key_info/></query></sirena>`)
	// response, err := s.ProxySendRequest(testRequestBytes, "key_info", "key_info")
	// if err != nil {
	// 	return err
	// }
	// logs.Log.Debugf("%s", response)
	return nil
}

// Request sends sirena XML request to sirena proxy
func (s *Repository) Request(requestBytes []byte, logAttributes map[string]string) ([]byte, error) {

	request, err := http.NewRequest("POST", "https://user:SUrPr5vj@sirena-proxy.dev.tmc24.io/", bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, err
	}

	response, err := s.p.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if logAttributes != nil {
		err = s.LogsPublisher.PublishLogs(logAttributes, requestBytes, responseBytes)
		if err != nil {
			return nil, err
		}
	}

	return responseBytes, nil
}
