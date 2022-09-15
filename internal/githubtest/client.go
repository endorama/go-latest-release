package githubtest

import (
	"log"
	"net/http"
	"path"
	"time"

	"github.com/dnaeon/go-vcr/v2/recorder"
)

// GetHttpClient instantiate a http.Client backed by a recorder.Recorder to be used in testing
// scenarios.
// NOTE: always remember to call (recorder.Recorder).Stop() in your test case like:
//
//	r, c := getHttpClient("a/fixture/path")
//	defer githubtest.StopOrFatal(r)
func GetHttpClient(fixture string) (*recorder.Recorder, *http.Client) {
	rec, err := recorder.New(path.Join("testdata", "fixtures", fixture))
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: rec,
		Timeout:   4 * time.Second,
	}

	return rec, client
}

func StopOrFatal(r *recorder.Recorder) {
	if err := r.Stop(); err != nil {
		log.Fatal(err)
	}
}
