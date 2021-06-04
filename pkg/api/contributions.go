package api

import (
	"context"

	"github.com/gleich/logoru"
	"github.com/shurcooL/githubv4"
)

type Query struct {
	Viewer struct {
		ContributionsCollection struct {
			ContributionCalendar struct {
				Colors      []string
				IsHalloween bool
				Weeks       []struct {
					ContributionDays []struct {
						Color string
					}
				}
			}
		}
	}
}

// Get contributions from the GitHub api
func Contributions(client *githubv4.Client) Query {
	var q Query
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		logoru.Error("Failed to get contributions from GitHub api;", err)
	}
	logoru.Info("Got contributions from GitHub API")
	return q
}
