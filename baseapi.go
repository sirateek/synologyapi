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

type ApiDetails struct {
	Host string
	Port int
	SSL  bool
}

type BaseApi struct {
	HttpClient    *http.Client
	ApiDetails    *ApiDetails
	ApiCredential *models.ApiCredential
	ApiInfo       models.ApiInfo
}

func (b *BaseApi) getHttpClient() *http.Client {
	if b.HttpClient == nil {
		b.HttpClient = &http.Client{}
	}
	return b.HttpClient
}

func (b *BaseApi) GetNewHttpRequest(httpMethod HttpMethod, api string) (*http.Request, error) {
	urlScheme := "%s://%s:%d/webapi/%s"
	protocol := "https"
	if !b.ApiDetails.SSL {
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
	url := fmt.Sprintf(urlScheme, protocol, b.ApiDetails.Host, b.ApiDetails.Port, value.Path)
	return http.NewRequest(string(httpMethod), url, nil)
}

func (b *BaseApi) SendRequest(req *http.Request, targetResponse any) error {
	// Http Client
	client := b.getHttpClient()

	if b.ApiCredential != nil {
		logrus.Trace("ApiCredential is not nil")
		sid := b.ApiCredential.GetSID()
		if sid != "" {
			req.URL.RawQuery += fmt.Sprint("&_sid=", sid)
			logrus.Trace("Adding _sid to the query")
		}
	}
	fmt.Println(req)
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
