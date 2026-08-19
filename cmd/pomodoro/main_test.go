package main

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name      string
		remaining time.Duration
		want      string
	}{
		{name: "開始前", remaining: 1 * time.Minute, want: "test: 01:00"},
		{name: "半分経過", remaining: time.Minute / 2, want: "test: 00:30"},
		{name: "終了", remaining: 0 * time.Minute, want: "test: 00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := format("test", tt.remaining, 0)

			if got != tt.want {
				t.Errorf("format() = %q, want %q", got, tt.want)
			}
		})
	}
}
