package httpcli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPMethod http method
type HTTPMethod string

const (
	// GET http get
	GET HTTPMethod = "GET"
	// POST http post
	POST HTTPMethod = "POST"
	// PUT http put
	PUT HTTPMethod = "PUT"
)

type httpBuilder struct {
	method      string
	url         string
	header      http.Header
	contentType http.Header
	body        io.Reader
	timeout     time.Duration
}

var defaultTimeout = time.Duration(30 * time.Second)

// NewBuilder http client builder
func NewBuilder(url string) *httpBuilder {
	return &httpBuilder{
		method:  string(GET),
		url:     url,
		header:  http.Header{},
		body:    nil,
		timeout: -1,
	}
}

// Method set http method
func (hb *httpBuilder) Method(method HTTPMethod) *httpBuilder {
	hb.method = string(method)
	return hb
}

// Timeout set http req timeout
func (hb *httpBuilder) Timeout(duration time.Duration) *httpBuilder {
	hb.timeout = duration
	return hb
}

func (hb *httpBuilder) AddHeader(key, value string) *httpBuilder {
	hb.header.Add(key, value)
	return hb
}

func (hb *httpBuilder) SetContentType(contentType string) *httpBuilder {
	hb.contentType = http.Header{}
	hb.contentType.Add("Content-Type", contentType)
	return hb
}

func (hb *httpBuilder) Body(body io.Reader) *httpBuilder {
	hb.body = body
	return hb
}

func (hb *httpBuilder) MapBody(params map[string]interface{}) *httpBuilder {
	bs, _ := json.Marshal(params)
	return hb.Body(strings.NewReader(string(bs)))
}

func (hb *httpBuilder) Do() (*http.Response, error) {
	return hb.DoWithProxy("")
}

func (hb *httpBuilder) DoWithProxy(proxyUrl string) (*http.Response, error) {
	if hb.timeout == -1 {
		hb.timeout = defaultTimeout
	}

	client := &http.Client{
		Timeout: hb.timeout,
	}
	if "" != proxyUrl {
		proxyURL, err := url.Parse(proxyUrl)
		if nil != err {
			fmt.Println("parse proxy url err,", err)
			return nil, err
		}
		proxy := http.ProxyURL(proxyURL)
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: proxy,
			},
		}
	}

	//build request
	request, err := http.NewRequest(hb.method, hb.url, hb.body)
	if nil != err {
		fmt.Println("new request err,", err)
		return nil, err
	}

	//add header
	for k, vals := range hb.header {
		for _, v := range vals {
			request.Header.Add(k, v)
		}
	}

	//set contentType
	for k, vals := range hb.contentType {
		request.Header.Del("Content-Type")
		request.Header.Add(k, vals[0])
	}

	return client.Do(request)
}
