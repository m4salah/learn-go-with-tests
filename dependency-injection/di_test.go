package main

import (
	"bytes"
	"testing"
)

func TestGreat(t *testing.T) {
	buffer := bytes.Buffer{}
	Great(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
