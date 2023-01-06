package main

import (
	"fmt"
	"net/http"

	"github.com/sirateek/synology_api_go/models"
	"github.com/sirateek/synology_api_go/synologyapi"
)

func main() {
	baseApi := synologyapi.BaseApi{
		HttpClient: &http.Client{},
		SSL:        true,
		Host:       "nas.sirateek.dev",
		Port:       5001,
	}

	auth := synologyapi.NewAuthenticate(baseApi)
	sid, err := auth.Login(models.AuthenticateRequest{
		Account: "tester",
		Passwd:  "tester1234",
	})

	fmt.Println("SID: ", sid, " ERR: ", err)
}
