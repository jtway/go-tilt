// Package tilt provides an interface to Tilt Bluetooth devices
package tilt

import (
	"encoding/hex"
	"log"
	"math"

	"github.com/pkg/errors"
)

// tiltIdentifier iBeacon identifier portion (4c000215) as well as Tilt specific uuid preamble (a495)
const tiltIdentifier = "4c000215a495"

// ErrNotTilt - the BLE device does not match anything in tiltType
var ErrNotTilt = errors.New("Not a Tilt iBeacon")

// Colour of the Tilt
type Colour string

var tiltType = map[string]Colour{
	"a495bb10c5b14b44b5121370f02d74de": "Red",
	"a495bb20c5b14b44b5121370f02d74de": "Green",
	"a495bb30c5b14b44b5121370f02d74de": "Black",
	"a495bb40c5b14b44b5121370f02d74de": "Purple",
	"a495bb50c5b14b44b5121370f02d74de": "Orange",
	"a495bb60c5b14b44b5121370f02d74de": "Blue",
	"a495bb70c5b14b44b5121370f02d74de": "Yellow",
	"a495bb80c5b14b44b5121370f02d74de": "Pink",
}

// Tilt struct
type Tilt struct {
	col  Colour
	temp uint16
	sg   uint16
}

// NewTilt returns a Tilt from an iBeacon
func NewTilt(b *IBeacon) (t Tilt, err error) {
	if col, ok := tiltType[b.UUID]; ok {
		t = Tilt{col: col, temp: b.Major, sg: b.Minor}
		return
	}
	err = ErrNotTilt
	return
}

// IsTilt tests if the data is from a Tilt
func IsTilt(d []byte) bool {
	if len(d) >= 25 && hex.EncodeToString(d)[0:12] == tiltIdentifier {
		return true
	}
	return false
}

func (t *Tilt) Celsius() float64 {
	return math.Round(float64(t.temp-32)/1.8*100) / 100
}

func (t *Tilt) Fahrenheit() uint16 {
	return t.temp
}

func (t *Tilt) Gravity() float64 {
	return float64(t.sg) / 1000
}

func (t *Tilt) Colour() Colour {
	return t.col
}

func (t *Tilt) Print() {
	log.Printf("Tilt: %v", t.Colour())
	log.Printf("Fahrenheit: %v\n", t.Fahrenheit())
	log.Printf("Specific Gravity: %v\n", t.Gravity())
	log.Printf("Celsius: %v\n", t.Celsius())
}
