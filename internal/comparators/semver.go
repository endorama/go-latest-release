package comparators

import (
	"fmt"

	"github.com/endorama/go-latest-release/internal/fetchers"
	version "github.com/hashicorp/go-version"
)

func NewSemver(current, new string) (*SemVer, error) {
	v1, err := version.NewSemver(current)
	if err != nil {
		return &SemVer{}, fmt.Errorf("cannot parse current version semver: %w", err)
	}

	v2, err := version.NewSemver(new)
	if err != nil {
		return &SemVer{}, fmt.Errorf("cannot parse new version semver: %w", err)
	}

	return &SemVer{current: v1, new: v2}, nil
}

type SemVer struct {
	current *version.Version
	new     *version.Version
}

func (c SemVer) Current() fetchers.Release {
	return fetchers.Release{Version: c.current}
}

func (c SemVer) Outdated() bool {
	return c.current.LessThan(c.new)
}

func (c SemVer) Latest() bool {
	return !c.Outdated()
}

func (c SemVer) New() fetchers.Release {
	return fetchers.Release{Version: c.new}
}
