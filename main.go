package main

import (
	"github.com/Matt-Gleich/contrihat/pkg/api"
)

func main() {
	client := api.Auth()
	api.Contributions(client)
}
