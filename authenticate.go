package synologyapi

import (
	"fmt"
	"math/rand"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/sirateek/synologyapi/models"
	"github.com/sirupsen/logrus"
)

type authenticate struct {
	BaseApi
	Api string
}

type Authenticate interface {
	Login(credential *models.ApiCredential) (sid string, err error)
	ReAuthenticate() (string, error)
	Logout() error
}

func NewAuthenticate(baseApi *BaseApi) Authenticate {
	return &authenticate{
		BaseApi: *baseApi,
		Api:     "SYNO.API.Auth",
	}
}

func (a *authenticate) Login(credential *models.ApiCredential) (string, error) {
	if credential.Session == "" {
		credential.Session = fmt.Sprint(rand.Intn(100))
	}

	// Prepare the query params.
	value, err := query.Values(credential)
	if err != nil {
		return "", err
	}
	value.Add("api", a.Api)
	value.Add("method", "login")
	value.Add("version", "3")

	// Inject the param into the request.
	req, err := a.GetNewHttpRequest(GET, a.Api)
	if err != nil {
		return "", err
	}
	req.URL.RawQuery = value.Encode()

	var targetResponse models.Response[models.AuthenticateResponse]
	err = a.SendRequest(req, &targetResponse)
	logrus.Debug(targetResponse)
	if err != nil {
		return "", err
	}

	return targetResponse.Data.Sid, nil
}

func (a *authenticate) ReAuthenticate() (string, error) {
	return a.Login(a.BaseApi.ApiCredential)
}

func (a *authenticate) Logout() error {
	value := url.Values{}

	value.Add("api", a.Api)
	value.Add("method", "logout")
	value.Add("version", "1")
	value.Add("session", a.BaseApi.ApiCredential.Session)

	// Inject the param into the request.
	req, err := a.GetNewHttpRequest(GET, a.Api)
	if err != nil {
		return err
	}
	req.URL.RawQuery = value.Encode()

	var targetResponse models.Response[any]
	err = a.SendRequest(req, targetResponse)
	if err != nil {
		return err
	}
	return nil
}
