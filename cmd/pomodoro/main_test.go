package main

import (
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	tests := []struct {
		name string
		done int
		want string
	}{
		{name: "done 0", done: 0, want: "░░░░░░░░░░░░░░░░░░░░"},
		{name: "done 10", done: 10, want: "██████████░░░░░░░░░░"},
		{name: "done 20", done: 20, want: "████████████████████"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bar(tt.done)

			if got != tt.want {
				t.Errorf("bar(%d) = %q, want %q", tt.done, got, tt.want)
			}
		})
	}
}

func TestProgressBar(t *testing.T) {
	tests := []struct {
		name      string
		remaining time.Duration
		total     time.Duration
		want      string
	}{
		{name: "開始前", remaining: 25 * time.Minute, total: 25 * time.Minute, want: "░░░░░░░░░░░░░░░░░░░░"},
		{name: "半分経過", remaining: 25 * time.Minute / 2, total: 25 * time.Minute, want: "██████████░░░░░░░░░░"},
		{name: "終了", remaining: 0 * time.Minute, total: 25 * time.Minute, want: "████████████████████"},
		{name: "total が 0", remaining: 0, total: 0, want: "████████████████████"},
		{name: "remaining が total 超", remaining: 30 * time.Minute, total: 25 * time.Minute, want: "░░░░░░░░░░░░░░░░░░░░"},
		{name: "remaining がマイナス", remaining: -1 * time.Second, total: 25 * time.Minute, want: "████████████████████"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := progressBar(tt.remaining, tt.total)

			if got != tt.want {
				t.Errorf("progressBar() = %q, want %q", got, tt.want)
			}
		})
	}
}

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
