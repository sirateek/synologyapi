package synologyapi

import (
	"net/url"

	"github.com/sirateek/synologyapi/models"
)

type infoApi struct {
	BaseApi
	Api string
}

type InfoApi interface {
	GetApisInfo() (*models.Response[models.ApiInfo], error)
}

func NewInfo(baseApi BaseApi) InfoApi {
	return &infoApi{
		BaseApi: baseApi,
		Api:     "SYNO.API.Info",
	}
}

func (i *infoApi) GetApisInfo() (*models.Response[models.ApiInfo], error) {
	value := url.Values{}
	value.Add("method", "query")
	value.Add("query", "all")

	req, err := i.GetNewHttpRequest(GET, i.Api, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery += value.Encode()
	var targetResponse models.Response[models.ApiInfo]
	err = i.SendRequest(req, &targetResponse)
	if err != nil {
		return nil, err
	}
	return &targetResponse, nil
}
