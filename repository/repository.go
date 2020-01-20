package repository

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/tmconsulting/sirenaxml-golang-sdk/publisher"
)

type Config struct {
	ProxyURL string
}

//Storage stow storage container
type Repository struct {
	p             *http.Client
	LogsPublisher publisher.Publisher
	Config        Config
}

func NewStorage(logsPublisher publisher.Publisher, opts ...Option) (*Repository, error) {
	r := new(Repository)

	r.p = &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}

	for _, opt := range opts {
		opt(r)
	}

	r.LogsPublisher = logsPublisher

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

// Request sends sirena XML request to sirena proxy
func (r *Repository) Request(requestBytes []byte, logAttributes map[string]string) ([]byte, error) {

	request, err := http.NewRequest("POST", r.Client.Conn.ProxyURL, bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, err
	}

	response, err := r.p.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if logAttributes != nil {
		err = r.LogsPublisher.PublishLogs(logAttributes, requestBytes, responseBytes)
		if err != nil {
			return nil, err
		}
	}

	return responseBytes, nil
}
