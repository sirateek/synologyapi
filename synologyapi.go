package synologyapi

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

type synologyApi struct {
	authenticate   AuthenticateApi
	baseApi        BaseApi
	infoApi        InfoApi
	certificateApi CertificateApi
}

type SynologyApi interface {
	Authenticate() AuthenticateApi
	Info() InfoApi
	Certificate() CertificateApi
	SetApiCredentialState(credentialState *apiCredentialState)
}

var (
	ErrGetApiInfo error = errors.New("there was a problem while getting the api spec from your synology.")
)

// A Facade pattern, Every one should create this instance before usage.
func NewSynologyApi(apiEndpoint *ApiEndpoint) (SynologyApi, error) {
	api := &synologyApi{
		baseApi: BaseApi{
			HttpClient:  &http.Client{},
			ApiEndpoint: apiEndpoint,
		},
	}
	apiInfo, err := api.Info().GetApisInfo()
	if err != nil {
		logrus.Error("Get API Info error: ", err)
		return api, ErrGetApiInfo
	}
	api.baseApi.ApiInfo = apiInfo.Data
	return api, nil
}

func (s *synologyApi) SetApiCredentialState(credentialState *apiCredentialState) {
	s.baseApi.apiCredentialState = credentialState
}

func (s *synologyApi) Authenticate() AuthenticateApi {
	if s.authenticate == nil {
		s.authenticate = NewAuthenticate(&s.baseApi)
	}
	return s.authenticate
}

func (s *synologyApi) Info() InfoApi {
	if s.infoApi == nil {
		s.infoApi = NewInfo(s.baseApi)
	}
	return s.infoApi
}

func (s *synologyApi) Certificate() CertificateApi {
	if s.certificateApi == nil {
		s.certificateApi = NewCertificate(&s.baseApi)
	}
	return s.certificateApi
}
