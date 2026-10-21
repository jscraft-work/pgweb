package static

import (
	"io/fs"
	"regexp"
	"strings"
	"testing"
)

func TestEmbeddedBrandReferences(t *testing.T) {
	html, err := fs.ReadFile(assets, "index.html")
	if err != nil {
		t.Fatal(err)
	}
	refs := regexp.MustCompile(`(?:src|href)="(static/brand/[^"]+)"`).FindAllSubmatch(html, -1)
	if len(refs) == 0 {
		t.Fatal("no brand assets declared")
	}
	for _, ref := range refs {
		name := strings.TrimPrefix(string(ref[1]), "static/")
		if _, err := fs.ReadFile(assets, name); err != nil {
			t.Errorf("brand asset %s missing from production binary: %v", name, err)
		}
	}
}
