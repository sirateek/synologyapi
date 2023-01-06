package synologyapi

import (
	"net/http"
)

type synologyApi struct {
	authenticate Authenticate
	baseApi      BaseApi
}

type SynologyApi interface {
	Authenticate() Authenticate
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
