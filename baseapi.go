package synologyapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirateek/synologyapi/models"
	"github.com/sirupsen/logrus"
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

func (b *BaseApi) getHttpClient() *http.Client {
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

func (b *BaseApi) SendRequest(req *http.Request, targetResponse any) error {
	// Http Client
	client := b.getHttpClient()

	if b.ApiCredential != nil {
		sid := b.ApiCredential.GetSID()
		if sid != "" {
			req.URL.RawQuery += fmt.Sprint("&_sid=", sid)
			logrus.Trace("Adding _sid to the query")
		}
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	// Read all response
	resBodyByte, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if targetResponse == nil {
		return nil
	}
	err = json.Unmarshal(resBodyByte, targetResponse)
	return err
}
