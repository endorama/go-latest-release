package fetchers_test

import (
	"context"
	"testing"

	"github.com/endorama/go-latest-release/internal/fetchers"
	"github.com/endorama/go-latest-release/internal/githubtest"
	"github.com/stretchr/testify/require"
)

func TestGitHubTagFetcher(t *testing.T) {
	r, c := githubtest.GetHttpClient("github/getlatesttag")
	defer githubtest.StopOrFatal(r)

	gtf := fetchers.GitHubLatestTag{
		Client:     c,
		Owner:      "endorama",
		Repository: "2ami",
	}

	rel, err := gtf.Fetch(context.Background())
	require.Nil(t, err)

	require.Equal(t, "0.6.0", rel.String())
}
