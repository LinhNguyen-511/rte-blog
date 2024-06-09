package templates

import (
	_ "embed"
	"testing"

	"github.com/a-h/templ/generator/htmldiff"
)

//go:embed snapshots/post_head.html
var postHead string

func TestHead(t *testing.T) {
	t.Run("returns a head component", func(t *testing.T) {
		got := head()
		diff, err := htmldiff.Diff(got, postHead)

		if err != nil {
			t.Fatal(err)
		}
		if diff != "" {
			t.Error(diff)
		}
	})
}
