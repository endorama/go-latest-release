package fetchers_test

import (
	"context"
	"testing"

	"github.com/endorama/go-latest-release/internal/fetchers"
	"github.com/endorama/go-latest-release/internal/githubtest"
	"github.com/stretchr/testify/require"
)

func TestGitHubLatestCommit(t *testing.T) {
	r, c := githubtest.GetHttpClient("github/getcommits")
	defer githubtest.StopOrFatal(r)

	gtf := fetchers.GitHubLatestCommit{
		Client:     c,
		Owner:      owner,
		Repository: repo,
	}

	rel, err := gtf.Fetch(context.Background())
	require.Nil(t, err)

	require.Equal(t, "575f1ae99b037666b356a93aae06348d925b8fd5", rel.GitHubCommit.Tree.GetSHA())
	require.Equal(t, "2022-09-15 14:21:09 +0000 UTC", rel.GitHubCommit.Committer.GetDate().String())
}

func TestGitHubLatestCommit_errorCode(t *testing.T) {
	r, c := githubtest.GetHttpClient("github/getcommits-non200")
	defer githubtest.StopOrFatal(r)

	gtf := fetchers.GitHubLatestCommit{
		Client:     c,
		Owner:      owner,
		Repository: "repo-that-does-not-exist",
	}

	_, err := gtf.Fetch(context.Background())
	require.Error(t, err)
}
