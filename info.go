package synologyapi

import (
	"net/url"

	"github.com/sirateek/synologyapi/models"
)

type infoApi struct {
	BaseApi
	Api     string
	cgiPath string
}

type InfoApi interface {
	GetApisInfo() (*models.Response[models.ApiInfo], error)
}

func NewInfo(baseApi BaseApi) InfoApi {
	return &infoApi{
		BaseApi: baseApi,
		Api:     "SYNO.API.Info",
		cgiPath: "query.cgi",
	}
}

func (i *infoApi) GetApisInfo() (*models.Response[models.ApiInfo], error) {
	value := url.Values{}
	value.Add("api", i.Api)
	value.Add("method", "query")
	value.Add("version", "1")
	value.Add("query", "all")

	req, err := i.GetNewHttpRequest(GET, i.cgiPath)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = value.Encode()
	var targetResponse models.Response[models.ApiInfo]
	err = i.SendRequest(req, &targetResponse)
	if err != nil {
		return nil, err
	}
	return &targetResponse, nil
}
