package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	var buf bytes.Buffer

	if err := run(&buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}

	want := "Hello, Go CLI!\n"
	if got := buf.String(); got != want {
		t.Errorf("run() output = %q, want %q", got, want)
	}
}
