package main

import (
	"time"

	"github.com/Matt-Gleich/contrihat/pkg/api"
	"github.com/Matt-Gleich/contrihat/pkg/lights"
)

func main() {
	client := api.Auth()
	for {
		contributions := api.Contributions(client)
		lights.Set(contributions)
		time.Sleep(3 * time.Second)
	}
}
