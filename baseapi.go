package synologyapi

import (
	"fmt"
	"net/http"

	"github.com/sirateek/synologyapi/models"
)

type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
)

type ApiDetails struct {
	Host string
	Port int
	SSL  bool
}

type BaseApi struct {
	HttpClient    *http.Client
	ApiDetails    *ApiDetails
	ApiCredential *models.ApiCredential
}

func (b *BaseApi) GetHttpClient() *http.Client {
	if b.HttpClient == nil {
		b.HttpClient = &http.Client{}
	}
	return b.HttpClient
}

func (b *BaseApi) GetNewHttpRequest(httpMethod HttpMethod, cgiPath string) (*http.Request, error) {
	urlScheme := "%s://%s:%d/webapi/%s"
	protocol := "https"
	if !b.ApiDetails.SSL {
		protocol = "http"
	}
	url := fmt.Sprintf(urlScheme, protocol, b.ApiDetails.Host, b.ApiDetails.Port, cgiPath)
	return http.NewRequest(string(httpMethod), url, nil)
}
