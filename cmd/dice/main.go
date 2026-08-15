package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

// roll は 1〜6 のサイコロの目を返す。
// intn は rand.Intn と同じく [0,n) の値を返す関数。
func roll(intn func(n int) int) int {
	return intn(6) + 1
}

func run(w io.Writer, intn func(n int) int) error {
	_, err := fmt.Fprintln(w, roll(intn))

	return err
}

func main() {
	if err := run(os.Stdout, rand.Intn); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
