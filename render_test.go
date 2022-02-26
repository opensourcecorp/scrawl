package main

import (
	"os"
	"path/filepath"
	"testing"
)

// This test suuucks lol
func TestGetCSSContents(t *testing.T) {
	got := getCSSContent("notes")

	if len(got) == 0 {
		t.Errorf("Returned CSS data is of length zero")
	}
}

func TestRenderFilesOutput(t *testing.T) {
	setTestScrawldir("web")
	scrawldir := getScrawldir()
	renderDir := filepath.Join(scrawldir, "web")

	// ugly T in the middle of these dummy filenames because scrawl notes are
	// split on the T from the RFC3339 format
	want := []string{
		filepath.Join(renderDir, "notedateTnotetime-1.html"),
		filepath.Join(renderDir, "notedateTnotetime-2.html"),
	}

	if err := os.RemoveAll(renderDir); err != nil {
		t.Errorf("\ncould not remove old rendered test data dir\n%v", err)
	}
	render()

	files, err := os.ReadDir(renderDir)
	if err != nil {
		t.Errorf("\ncould not read files from %s\n%v", renderDir, err)
	}

	var got []string
	for _, file := range files {
		got = append(got, filepath.Join(scrawldir, "web", file.Name()))
	}

	cmpEqualWantGot(t, want, got)
}
