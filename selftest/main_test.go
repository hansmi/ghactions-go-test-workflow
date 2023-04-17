package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdd(t *testing.T) {
	got := add(100, 222)
	want := 322

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("add() diff (-want +got):\n%s", diff)
	}
}
