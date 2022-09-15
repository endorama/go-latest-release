package fetchers

import (
	"context"
	"fmt"
	"net/http"

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

func (g GitHubLatestTag) Fetch(ctx context.Context) (*version.Version, error) {
	client := github.NewClient(g.Client)

	rels, resp, err := client.Repositories.GetLatestRelease(ctx, g.Owner, g.Repository)
	if err != nil {
		return &version.Version{}, fmt.Errorf("cannot retrieve latest release: %w", err)
	}

	if resp.StatusCode != 200 {
		return &version.Version{}, fmt.Errorf("status code was not 200: status code %d", resp.StatusCode)
	}

	v, err := version.NewSemver(rels.GetTagName())
	if err != nil {
		return &version.Version{}, err
	}

	return v, err
}
