package fetchers

import (
	"context"
	"fmt"
	"net/http"

	latestrelease "github.com/endorama/go-latest-release"
	"github.com/google/go-github/github"
)

type GitHubLatestCommit struct {
	Client *http.Client

	Owner      string
	Repository string
}

func (g GitHubLatestCommit) Fetch(ctx context.Context) (latestrelease.Release, error) {
	client := github.NewClient(g.Client)

	commitInfo, resp, err := client.Repositories.ListCommits(ctx, g.Owner, g.Repository, &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	})
	if err != nil {
		return latestrelease.Release{}, fmt.Errorf("cannot retrive commit information: %w", err)
	}

	if resp.StatusCode != 200 {
		return latestrelease.Release{}, fmt.Errorf("status code was not 200: status code %d", resp.StatusCode)
	}

	return latestrelease.Release{GitHubCommit: commitInfo[0].Commit}, err
}
