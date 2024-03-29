package api

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gleich/contrihat/pkg/config"
	"github.com/gleich/logoru"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func Auth() *githubv4.Client {
	// Reading from pat.txt file
	b, err := ioutil.ReadFile(filepath.Join(config.GetFolderName(), "pat.txt"))
	if err != nil {
		logoru.Critical("Failed to read from pat.txt file")
		os.Exit(1)
	}

	// Creating client
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: strings.TrimSuffix(string(b), "\n")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return githubv4.NewClient(httpClient)
}
