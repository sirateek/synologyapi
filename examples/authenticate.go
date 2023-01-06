//go:build exclude

package examples

import (
	"fmt"

	"github.com/sirateek/synologyapi"
	"github.com/sirateek/synologyapi/models"
)

func examples() {
	api := synologyapi.NewSynologyApi(&synologyapi.ApiDetails{
		Host: "example.com",
		Port: 5001,
		SSL:  true,
	})

	sid, err := api.Authenticate().Login(models.ApiCredential{
		Account: "user",
		Passwd:  "password",
		Format:  "sid",
	})

	fmt.Println("SID: ", sid, "ERR: ", err)

	// You can start calling other apis without any futher interaction with Authenticate api
	// The `Login()` method just return the sid incase you need it.
}
