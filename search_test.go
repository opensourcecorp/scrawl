package main

import (
	"testing"
)

func TestSearchTagsAnyFlagEmptyString(t *testing.T) {
	setTestScrawldir("tags")

	want := tagMap{
		"testdata/tags/dev-todo.md":  []string{"dev", "todo"},
		"testdata/tags/dev-work.md":  []string{"dev", "work"},
		"testdata/tags/dev.md":       []string{"dev"},
		"testdata/tags/work-todo.md": []string{"work", "todo"},
	}

	got := searchTags(true, false, false, "")

	cmpEqualWantGot(t, want, got)
}

func TestSearchTagsAnyFlagOneTag(t *testing.T) {
	setTestScrawldir("tags")

	want := tagMap{
		"testdata/tags/dev-todo.md": []string{"dev", "todo"},
		"testdata/tags/dev-work.md": []string{"dev", "work"},
		"testdata/tags/dev.md":      []string{"dev"},
	}

	got := searchTags(true, false, false, "dev")

	cmpEqualWantGot(t, want, got)
}

func TestSearchTagsAnyFlagOneTagRegex(t *testing.T) {
	setTestScrawldir("tags")

	want := tagMap{
		"testdata/tags/dev-work.md":  []string{"dev", "work"},
		"testdata/tags/work-todo.md": []string{"work", "todo"},
	}

	got := searchTags(true, false, false, "wor.*")

	cmpEqualWantGot(t, want, got)
}

func TestSearchTagsUntagged(t *testing.T) {
	setTestScrawldir("tags")

	want := tagMap{
		"testdata/tags/untagged.md": []string{},
	}

	got := searchTags(true, false, true, "")

	cmpEqualWantGot(t, want, got)
}

func TestGetTagsFromFileContentsOneTag(t *testing.T) {
	setTestScrawldir("tags")

	contents := []string{
		"tags: dev",
		"",
		"this is an embedded test note, wow",
	}

	want := []string{"dev"}
	got := getTagsFromFileContents(contents)

	cmpEqualWantGot(t, want, got)
}

func TestGetTagsFromFileContentsMultipleTags(t *testing.T) {
	setTestScrawldir("tags")

	contents := []string{
		"tags: dev, work",
		"",
		"this is an embedded test note, wow",
	}

	want := []string{"dev", "work"}
	got := getTagsFromFileContents(contents)

	cmpEqualWantGot(t, want, got)
}
