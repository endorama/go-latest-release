package fetchers

import (
	"context"
	"fmt"
	"net/http"

	latestrelease "github.com/endorama/go-latest-release"
	"github.com/google/go-github/github"
	version "github.com/hashicorp/go-version"
)

type GitHubLatestTag struct {
	// The HTTP client to use for interacting with GitHub APIs
	Client *http.Client

	// Owner is the repository owner name
	Owner string
	// The repository name
	Repository string
}

func (g GitHubLatestTag) Fetch(ctx context.Context) (latestrelease.Release, error) {
	client := github.NewClient(g.Client)
	// client.BaseURL, _ = url.Parse(opts.URL)

	rels, resp, err := client.Repositories.GetLatestRelease(ctx, g.Owner, g.Repository)
	if err != nil {
		return latestrelease.Release{}, fmt.Errorf("cannot retrieve latest release: %w", err)
	}

	if resp.StatusCode != 200 {
		return latestrelease.Release{}, fmt.Errorf("status code was not 200: status code %d", resp.StatusCode)
	}

	v, err := version.NewSemver(rels.GetTagName())
	if err != nil {
		return latestrelease.Release{}, err
	}

	return latestrelease.Release{Version: v}, err
}
