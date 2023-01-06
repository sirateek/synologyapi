package synologyapi

import (
	"fmt"
	"net/http"
)

type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
)

type BaseApi struct {
	HttpClient *http.Client
	cgiPath    string
	Host       string
	Port       int
	SSL        bool
	SID        string
}

func (b *BaseApi) GetHttpClient() *http.Client {
	if b.HttpClient == nil {
		b.HttpClient = &http.Client{}
	}
	return b.HttpClient
}

func (b *BaseApi) GetNewHttpRequest(httpMethod HttpMethod) (*http.Request, error) {
	urlScheme := "%s://%s:%d/webapi/%s"
	protocol := "https"
	if !b.SSL {
		protocol = "http"
	}
	url := fmt.Sprintf(urlScheme, protocol, b.Host, b.Port, b.cgiPath)
	return http.NewRequest(string(httpMethod), url, nil)
}
