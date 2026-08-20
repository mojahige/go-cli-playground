package main

import (
	"fmt"
	"strings"
	"time"
)

const barWidth = 20

func bar(done int) string {
	return strings.Repeat("█", done) + strings.Repeat("░", barWidth-done)
}

func progressBar(remaining, total time.Duration) string {
	if total <= 0 || remaining <= 0 {
		return bar(barWidth)
	}

	if remaining >= total {
		return bar(0)
	}

	elapsed := total - remaining
	ratio := float64(elapsed) / float64(total)
	done := int(float64(barWidth) * ratio)

	return bar(done)
}

func format(label string, remaining, total time.Duration) string {
	m := int(remaining / time.Minute)
	s := int((remaining % time.Minute) / time.Second)

	return fmt.Sprintf("%s: %02d:%02d", label, m, s)
}
