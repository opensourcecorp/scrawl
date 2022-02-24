// Helpers for tests across the package
package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func cmpEqualWantGot(t *testing.T, got, want interface{}) {
	if !cmp.Equal(want, got) {
		t.Errorf("\nwant: %v\ngot:%v\n", want, got)
	}
}
