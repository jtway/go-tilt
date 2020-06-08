package main

import (
	"time"

	"github.com/alexhowarth/go-tilt"
)

func main() {

	s := tilt.NewScanner()
	s.Scan(20 * time.Second)

	for _, t := range s.Tilts() {
		t.Print()
	}
}
