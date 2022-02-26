package main

// Helpers for tests across the package

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func setTestScrawldir(testGroup string) {
	err := os.Setenv("SCRAWLDIR", "./testdata/"+testGroup)
	if err != nil {
		panic(err)
	}
}

func cmpEqualWantGot(t *testing.T, want, got interface{}) {
	if !cmp.Equal(want, got) {
		t.Errorf("\nwant: %v\ngot:  %v\n", want, got)
	}
}
