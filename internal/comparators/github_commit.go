package comparators

import (
	latestrelease "github.com/endorama/go-latest-release"
	"github.com/google/go-github/github"
)

func NewGitHubCommit(current, new latestrelease.Release) (*GitHubCommit, error) {
	return &GitHubCommit{current: current.GitHubCommit, new: new.GitHubCommit}, nil
}

type GitHubCommit struct {
	current *github.Commit
	new     *github.Commit
}

func (c GitHubCommit) Current() latestrelease.Release {
	return latestrelease.Release{GitHubCommit: c.current}
}

func (c GitHubCommit) Outdated() bool {
	return c.current.Committer.GetDate().Before(c.new.Committer.GetDate())
}

func (c GitHubCommit) Latest() bool {
	return !c.Outdated()
}

func (c GitHubCommit) New() latestrelease.Release {
	return latestrelease.Release{GitHubCommit: c.new}
}
