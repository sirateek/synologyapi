package synologyapi

import (
	"net/http"
)

type synologyApi struct {
	authenticate Authenticate
	baseApi      BaseApi
	infoApi      InfoApi
}

type SynologyApi interface {
	Authenticate() Authenticate
	Info() InfoApi
}

// A Facade pattern, Every one should create this instance before usage.
func NewSynologyApi(apiDetails *ApiDetails) SynologyApi {
	api := &synologyApi{
		baseApi: BaseApi{
			HttpClient: &http.Client{},
			ApiDetails: apiDetails,
		},
	}
	return api
}

func (s *synologyApi) Authenticate() Authenticate {
	if s.authenticate == nil {
		s.authenticate = NewAuthenticate(s.baseApi)
	}
	return s.authenticate
}

func (s *synologyApi) Info() InfoApi {
	if s.infoApi == nil {
		s.infoApi = NewInfo(s.baseApi)
	}
	return s.infoApi
}
