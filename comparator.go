package latestrelease

import (
	"github.com/google/go-github/github"
	version "github.com/hashicorp/go-version"
)

type Comparator interface {
	CurrentVersion() Release
	IsOutdated() bool
	IsLatest() bool
	New() Release
}

type Release struct {
	GitHubCommit *github.Commit
	Version      *version.Version
}

func (r Release) Get() string {
	switch {
	case r.GitHubCommit != nil:
		return r.GitHubCommit.GetSHA()
	case r.Version != nil:
		return r.Version.String()
	default:
		panic("no commit or version")
	}
}

type ReleaseInfo interface {
	Get() string

	GetGitHubCommit() *github.Commit
	GetVersion() *version.Version
}
