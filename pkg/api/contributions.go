package api

import (
	"context"
	"fmt"

	"github.com/Matt-Gleich/logoru"
	"github.com/shurcooL/githubv4"
)

// Get contributions from the GitHub api
func Contributions(client *githubv4.Client) {
	var q struct {
		Viewer struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					Colors []string
					Weeks  []struct {
						ContributionDays []struct {
							Color string
						}
					}
				}
			}
		}
	}
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		logoru.Error("Failed to get contributions from GitHub api;", err)
	}
	fmt.Printf("%#v", q)
}
