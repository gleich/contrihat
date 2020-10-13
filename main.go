package main

import (
	"time"

	"github.com/Matt-Gleich/contrihat/pkg/api"
	"github.com/Matt-Gleich/contrihat/pkg/lights"
)

func main() {
	client := api.Auth()
	firstRun := true
	for {
		contributions := api.Contributions(client)
		lights.Set(contributions, firstRun)
		time.Sleep(5 * time.Second)
		firstRun = false
	}
}
