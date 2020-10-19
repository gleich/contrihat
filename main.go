package main

import (
	"math/rand"
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

		sleepSeconds := time.Duration(int64(time.Second) * int64(rand.Intn(configuration.UpdateInterval)))
		time.Sleep(sleepSeconds)
		firstRun = false
	}
}
