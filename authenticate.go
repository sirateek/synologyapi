package synologyapi

import (
	"fmt"
	"math/rand"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/sirateek/synologyapi/models"
	"github.com/sirupsen/logrus"
)

type authenticateApi struct {
	baseApi *BaseApi
	Api     string
}

type AuthenticateApi interface {
	Login(credential models.ApiCredential) (apiCredentialState, error)
	Logout(credentialState *apiCredentialState) error
}

type apiCredentialState struct {
	account     string `url:"account"`
	passwd      string `url:"passwd"`
	session     string `url:"session,omitempty"`
	format      string `url:"format,omitempty"`
	otpCode     string `url:"otp_code,omitempty"`
	sid         string
	isSignedOut bool
}

func NewAuthenticate(baseApi *BaseApi) AuthenticateApi {
	return &authenticateApi{
		baseApi: baseApi,
		Api:     "SYNO.API.Auth",
	}
}

func (a *authenticateApi) Login(credential models.ApiCredential) (state apiCredentialState, err error) {
	if credential.Session == "" {
		credential.Session = fmt.Sprint(rand.Intn(100))
	}

	var sid string

	// Prepare the query params.
	value, err := query.Values(credential)
	if err != nil {
		return state, err
	}
	value.Add("api", a.Api)
	value.Add("method", "login")
	value.Add("version", "3")

	// Inject the param into the request.
	req, err := a.baseApi.GetNewHttpRequest(GET, a.Api)
	if err != nil {
		return state, err
	}
	req.URL.RawQuery = value.Encode()

	var targetResponse models.Response[models.AuthenticateResponse]
	err = a.baseApi.SendRequest(req, &targetResponse)
	logrus.Debug(targetResponse)
	if err != nil {
		return state, err
	}
	sid = targetResponse.Data.Sid

	return apiCredentialState{
		account:     credential.Account,
		passwd:      credential.Passwd,
		format:      credential.Format,
		otpCode:     credential.OtpCode,
		session:     credential.Session,
		sid:         sid,
		isSignedOut: false,
	}, nil
}

func (a *authenticateApi) Logout(credentialState *apiCredentialState) error {
	value := url.Values{}

	value.Add("api", a.Api)
	value.Add("method", "logout")
	value.Add("version", "1")
	value.Add("session", credentialState.session)

	// Inject the param into the request.
	req, err := a.baseApi.GetNewHttpRequest(GET, a.Api)
	if err != nil {
		return err
	}
	req.URL.RawQuery = value.Encode()

	var targetResponse models.Response[any]
	err = a.baseApi.SendRequest(req, &targetResponse)
	if err != nil {
		return err
	}
	credentialState.isSignedOut = true
	return nil
}
