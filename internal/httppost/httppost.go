package httppost

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"omnis-client/internal"
	"time"
)

type IHTTPClient interface {
	Post(IP string, contentType string, body io.Reader) (*http.Response, error)
}

type HTTPClient struct {
	Transport http.RoundTripper
	Timeout   time.Duration
}

type Fetcher struct {
	Netclient IHTTPClient
}

func NewFetcher(timeout int) *Fetcher {
	tr := &http.Transport{}

	var netClient = &http.Client{
		Transport: tr,
		Timeout:   time.Second * time.Duration(timeout),
	}
	return &Fetcher{
		Netclient: netClient,
	}
}

func (f Fetcher) Post(IP string, body []byte) (*internal.HTTPResponse, error) {

	resp, err := f.Netclient.Post(IP, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(respBody)

	var r = &internal.HTTPResponse{
		Body:       bodyString,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
	}
	return r, nil
}
