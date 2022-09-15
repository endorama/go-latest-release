package fetchers

import (
	"context"

	"github.com/hashicorp/go-version"
)

func New(s ReleaseFetcher) Fetcher {
	f := Fetcher{}
	f.SetStrategy(s)
	return f
}

type ReleaseFetcher interface {
	Fetch(ctx context.Context) (*version.Version, error)
}

type Fetcher struct {
	strategy ReleaseFetcher
}

func (f *Fetcher) SetStrategy(s ReleaseFetcher) {
	f.strategy = s
}

func (f Fetcher) Execute(ctx context.Context) (*version.Version, error) {
	return f.strategy.Fetch(ctx)
}
