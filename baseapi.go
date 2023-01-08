package synologyapi

import (
	"encoding/json"
	"errors"
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

type ApiEndpoint struct {
	Host string
	Port int
	SSL  bool
}

type BaseApi struct {
	HttpClient         *http.Client
	ApiEndpoint        *ApiEndpoint
	apiCredentialState *apiCredentialState
	ApiInfo            models.ApiInfo
}

func (b *BaseApi) SetApiCredentialState(credentialState *apiCredentialState) {
	b.apiCredentialState = credentialState
}

func (b *BaseApi) getHttpClient() *http.Client {
	if b.HttpClient == nil {
		b.HttpClient = &http.Client{}
	}
	return b.HttpClient
}

func (b *BaseApi) GetNewHttpRequest(httpMethod HttpMethod, api string, body io.Reader) (*http.Request, error) {
	urlScheme := "%s://%s:%d/webapi/%s"
	protocol := "https"
	if !b.ApiEndpoint.SSL {
		protocol = "http"
	}

	value, ok := b.ApiInfo[api]
	if !ok {
		if api != "SYNO.API.Info" {
			return nil, errors.New("api is not found")
		}
		value = models.ApiDetails{
			Path:       "entry.cgi",
			MinVersion: 1,
			Maxversion: 1,
		}
	}

	url := fmt.Sprintf(urlScheme, protocol, b.ApiEndpoint.Host, b.ApiEndpoint.Port, value.Path)
	req, err := http.NewRequest(string(httpMethod), url, body)

	if err != nil {
		return req, err
	}

	req.URL.RawQuery = fmt.Sprint("version=", value.Maxversion, "&")
	req.URL.RawQuery += fmt.Sprint("api=", api, "&")
	return req, nil
}

func (b *BaseApi) SendRequest(req *http.Request, targetResponse any) error {
	// Http Client
	client := b.getHttpClient()

	if b.apiCredentialState != nil {
		logrus.Trace("ApiCredential is not nil and is not signed out")
		sid := b.apiCredentialState.sid
		if sid != "" {
			req.URL.RawQuery += fmt.Sprint("&_sid=", sid)
			logrus.Trace("Adding _sid to the query")
		}
	}
	logrus.Debug("Request: ", req)
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	// Read all response
	resBodyByte, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	logrus.Debug("Response: ", response)
	logrus.Debug("Response Body: ", string(resBodyByte))

	if targetResponse == nil {
		return nil
	}
	err = json.Unmarshal(resBodyByte, targetResponse)
	return err
}
