package proxy

import (
	"fmt"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
)

// Memory storage keeps data in memory
type storage struct {
	proxyPath string
	r         *resty.Client
}

func NewStorage(proxyPath string, l logs.LogWriter, debug bool) *storage {
	s := new(storage)

	s.proxyPath = proxyPath
	s.r = resty.New()
	s.r.SetDebug(debug)
	s.r.SetLogger(l)
	s.r.SetRedirectPolicy(resty.FlexibleRedirectPolicy(5))

	return s
}

func (s *storage) SendRawRequest(req []byte) ([]byte, error) {
	return s.sendMsg(req)
}

func (s *storage) sendMsg(req []byte) ([]byte, error) {
	resp, err := s.r.R().SetBody(req).Post(s.proxyPath)
	if err != nil || resp.StatusCode() != 200 {
		if err == nil {
			return nil, errors.New(fmt.Sprintf("proxy request error: %v", string(resp.Body())))
		}
		return nil, err
	}

	return resp.Body(), nil
}
