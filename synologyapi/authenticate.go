package synologyapi

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/go-querystring/query"
	"github.com/sirateek/synology_api_go/models"
)

type authenticate struct {
	BaseApi
	Api string
}

type Authenticate interface {
	Login(credential models.AuthenticateRequest) (sid string, err error)
}

func NewAuthenticate(baseApi BaseApi) Authenticate {
	baseApi.cgiPath = "auth.cgi"
	return &authenticate{
		BaseApi: baseApi,
		Api:     "SYNO.API.Auth",
	}
}

func (a *authenticate) Login(credential models.AuthenticateRequest) (string, error) {
	// Prepare the query params.
	value, err := query.Values(credential)
	if err != nil {
		return "", err
	}
	value.Add("api", a.Api)
	value.Add("method", "login")
	value.Add("version", "3")

	// Inject the param into the request.
	req, err := a.GetNewHttpRequest(GET)
	if err != nil {
		return "", err
	}
	req.URL.RawQuery = value.Encode()

	response, err := a.GetHttpClient().Do(req)
	if err != nil {
		return "", err
	}

	var objmap map[string]json.RawMessage

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(body, &objmap)
	fmt.Println(string(objmap["success"]))

	return "", nil
}
