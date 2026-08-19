package main

import (
	"fmt"
	"time"
)

func format(label string, remaining, _ time.Duration) string {
	m := int(remaining / time.Minute)
	s := int((remaining % time.Minute) / time.Second)

	return fmt.Sprintf("%s: %02d:%02d", label, m, s)
}
