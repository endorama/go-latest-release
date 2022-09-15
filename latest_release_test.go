package latestrelease_test

import (
	"testing"

	"github.com/endorama/go-latest-release/internal/fetchers"
)

func TestAPI(t *testing.T) {
	current := "0.1.0"
	latest := fetchers.New(fetchers.GitHubLatestCommit{})
}
