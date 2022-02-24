// Helpers for tests across the package
package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func setTestScrawldir() {
	err := os.Setenv("SCRAWLDIR", "./testdata")
	if err != nil {
		panic(err)
	}
}

func cmpEqualWantGot(t *testing.T, got, want interface{}) {
	if !cmp.Equal(want, got) {
		t.Errorf("\nwant: %v\ngot:%v\n", want, got)
	}
}
