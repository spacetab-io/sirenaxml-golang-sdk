package socket

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	sirenaXML "github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/socket/client"
)

type storage struct {
	c *client.Channel
	p *http.Client
}

func NewClient(sc *sirenaXML.Config, l logs.LogWriter) (*storage, error) {
	c, err := client.NewChannel(sc, l)
	if err != nil {
		return nil, errors.Wrap(err, "sirena client init error")
	}
	return &storage{c: c}, nil
}

func (s *storage) SendRawRequest(req []byte) ([]byte, error) {
	return s.c.SendMsg(req)
}

// Request sends sirena XML request to sirena proxy
func (s *storage) Request(requestBytes []byte, logAttributes map[string]string) ([]byte, error) {

	request, err := http.NewRequest("POST", "", bytes.NewBuffer(requestBytes))
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

	return responseBytes, nil
}
