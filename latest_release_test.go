package latestrelease_test

import (
	"context"
	"testing"

	"github.com/endorama/go-latest-release/internal/fetchers"
	"github.com/endorama/go-latest-release/internal/githubtest"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
)

func ExampleVersion(t *testing.T) {
	// use a "real" HTTP client
	r, c := githubtest.GetHttpClient("github/getlatesttag")
	defer githubtest.StopOrFatal(r)

	current, err := version.NewSemver("0.5.0")
	assert.NoError(t, err)

	fetcher := fetchers.New(fetchers.GitHubLatestTag{Client: c, Owner: "endorama", Repository: "2ami"})
	latest, err := fetcher.Execute(context.Background())
	assert.NoError(t, err)

	assert.True(t, latest.GreaterThan(current))
	assert.Equal(t, "0.6.0", latest.String())
}
