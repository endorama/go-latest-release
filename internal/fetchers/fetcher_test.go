package fetchers_test

import (
	"context"
	"testing"

	"github.com/endorama/go-latest-release/internal/fetchers"
	"github.com/endorama/go-latest-release/internal/githubtest"
	"github.com/stretchr/testify/require"
)

const owner = "endorama"
const repo = "2ami"

func TestFetcher_GitHubTag(t *testing.T) {
	r, c := githubtest.GetHttpClient("github/getlatesttag")
	defer githubtest.StopOrFatal(r)

	gtf := fetchers.GitHubLatestTag{
		Client:     c,
		Owner:      owner,
		Repository: repo,
	}

	f := fetchers.Fetcher{}
	f.SetStrategy(gtf)

	rel, err := f.Execute(context.Background())
	require.Nil(t, err)
	require.Equal(t, "0.6.0", rel.Version.String())
}
