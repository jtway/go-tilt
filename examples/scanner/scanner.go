package main

import (
	"time"

	"github.com/jtway/go-tilt"
)

func main() {

	s := tilt.NewScanner()
	s.Scan(10 * time.Second)

	for _, t := range s.Tilts() {
		t.Print()
	}
	// Scane again to test
	s.Scan(10 * time.Second)

	for _, t := range s.Tilts() {
		t.Print()
	}
}
