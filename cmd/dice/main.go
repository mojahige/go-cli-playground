package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
)

const prompt = "Enter でサイコロを振るよ🎲 (Ctrl+D で終了)"
const clearScreen = "\033[H\033[2J"

// faces はサイコロの目の絵。index が出目に対応する（0 は未使用）。
var faces = [7]string{
	1: `┌───────┐
│       │
│   ●   │
│       │
└───────┘`,
	2: `┌───────┐
│ ●     │
│       │
│     ● │
└───────┘`,
	3: `┌───────┐
│ ●     │
│   ●   │
│     ● │
└───────┘`,
	4: `┌───────┐
│ ●   ● │
│       │
│ ●   ● │
└───────┘`,
	5: `┌───────┐
│ ●   ● │
│   ●   │
│ ●   ● │
└───────┘`,
	6: `┌───────┐
│ ●   ● │
│ ●   ● │
│ ●   ● │
└───────┘`,
}

// roll は 1〜6 のサイコロの目を返す。
// intn は rand.Intn と同じく [0,n) の値を返す関数。
func roll(intn func(n int) int) int {
	return intn(6) + 1
}

func run(r io.Reader, w io.Writer, intn func(n int) int) error {
	fmt.Fprintln(w, prompt)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fmt.Fprint(w, clearScreen)
		fmt.Fprintln(w, prompt)
		fmt.Fprintln(w, faces[roll(intn)])
	}

	return scanner.Err()
}

func main() {
	if err := run(os.Stdin, os.Stdout, rand.Intn); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
