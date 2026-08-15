package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRoll(t *testing.T) {
	tests := []struct {
		name string
		intn int
		want int
	}{
		{name: "最小値", intn: 0, want: 1},
		{name: "最大値", intn: 5, want: 6},
		{name: "途中", intn: 2, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intn := func(n int) int { return tt.intn }

			if got := roll(intn); got != tt.want {
				t.Errorf("roll() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFaces(t *testing.T) {
	for n := 1; n <= 6; n++ {
		if got := strings.Count(faces[n], "●"); got != n {
			t.Errorf("faces[%d] の目の数 = %d, want %d", n, got, n)
		}
	}
}

func TestRun(t *testing.T) {
	var buf bytes.Buffer
	in := strings.NewReader("\n\n\n") // Enter 3回
	intn := func(n int) int { return 2 }

	if err := run(in, &buf, intn); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}

	want := prompt + "\n" + strings.Repeat(clearScreen+prompt+"\n"+faces[3]+"\n", 3)
	if got := buf.String(); got != want {
		t.Errorf("run() output = %q, want %q", got, want)
	}
}
