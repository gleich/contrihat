package main

import (
	"time"

	"github.com/Matt-Gleich/contrihat/pkg/api"
	"github.com/Matt-Gleich/contrihat/pkg/config"
	"github.com/Matt-Gleich/contrihat/pkg/lights"
)

func main() {
	client := api.Auth()
	firstRun := true
	for {
		configuration := config.ReadGeneral()
		contributions := api.Contributions(client)
		lights.Set(contributions, firstRun, configuration)

		time.Sleep(3 * time.Second)
		firstRun = false
	}
}
