package synologyapi

import (
	"net/http"

	"github.com/sirateek/synologyapi/models"
)

type synologyApi struct {
	authenticate   Authenticate
	baseApi        BaseApi
	infoApi        InfoApi
	certificateApi CertificateApi
}

type SynologyApi interface {
	Authenticate() Authenticate
	Info() InfoApi
	Certificate() CertificateApi
}

// A Facade pattern, Every one should create this instance before usage.
func NewSynologyApi(apiDetails *ApiDetails, credential *models.ApiCredential) (SynologyApi, error) {
	api := &synologyApi{
		baseApi: BaseApi{
			HttpClient: &http.Client{},
			ApiDetails: apiDetails,
		},
	}
	apiInfo, err := api.Info().GetApisInfo()
	if err != nil {
		return api, err
	}
	api.baseApi.ApiInfo = apiInfo.Data

	_, err = api.Authenticate().Login(credential)
	if err != nil {
		return api, err
	}
	api.baseApi.ApiCredential = credential
	return api, nil
}

func (s *synologyApi) Authenticate() Authenticate {
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
