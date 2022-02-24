package main

import (
	"os"
	"testing"
)

func TestMakeNoteFilepathLength(t *testing.T) {
	setTestScrawldir()

	scrawldirLength := len(os.Getenv("SCRAWLDIR")) + 1 // directory length plus separator
	want := scrawldirLength + 19 + 3                   // timestamp length plus extension
	got := len(constructNoteFilepath())

	if want != got {
		t.Errorf("\nwant: %v\ngot: %v\n", want, got)
	}
}
