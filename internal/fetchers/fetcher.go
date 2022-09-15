package fetchers

import (
	"context"

	latestrelease "github.com/endorama/go-latest-release"
)

func New(s ReleaseFetcher) Fetcher {
	f := Fetcher{}
	f.SetStrategy(s)
	return f
}

type ReleaseFetcher interface {
	Fetch(ctx context.Context) (latestrelease.Release, error)
}

type Fetcher struct {
	strategy ReleaseFetcher
}

func (f *Fetcher) SetStrategy(s ReleaseFetcher) {
	f.strategy = s
}

func (f Fetcher) Execute(ctx context.Context) (latestrelease.Release, error) {
	return f.strategy.Fetch(ctx)
}
