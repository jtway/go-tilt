package tilt_test

import (
	"errors"
	"testing"

	"github.com/alexhowarth/go-tilt"
)

func TestNewTilt(t *testing.T) {
	tt := []struct {
		name       string
		ibeacon    *tilt.IBeacon
		colour     tilt.Colour
		celsius    float64
		fahrenheit uint16
		gravity    float64
		err        error
	}{
		{
			name:       "Red iBeacon",
			ibeacon:    &tilt.IBeacon{UUID: "a495bb10c5b14b44b5121370f02d74de", Major: 70, Minor: 1035},
			colour:     tilt.Colour("Red"),
			fahrenheit: 70,
			celsius:    21.11,
			gravity:    1.035,
			err:        nil,
		},
		{
			name:       "Black iBeacon",
			ibeacon:    &tilt.IBeacon{UUID: "a495bb30c5b14b44b5121370f02d74de", Major: 69, Minor: 1065},
			colour:     tilt.Colour("Black"),
			fahrenheit: 69,
			celsius:    20.56,
			gravity:    1.065,
			err:        nil,
		},
		{
			name:    "Not an iBeacon",
			ibeacon: &tilt.IBeacon{UUID: "a495bb99c5b14b44b5121370f02d74de", Major: 1, Minor: 2},
			err:     tilt.ErrNotTilt,
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			got, err := tilt.NewTilt(tc.ibeacon)

			if tc.err != nil {
				// expecting an error
				if !errors.Is(err, tc.err) {
					t.Fatalf("Expected '%v' error, got '%v' error", tc.err, err)
				}
				return
			}

			if got.Colour() != tc.colour {
				t.Errorf("Expected %v, got %v", tc.colour, got.Colour())
			}
			if got.Celsius() != tc.celsius {
				t.Errorf("Expected %v, got %v", tc.celsius, got.Celsius())
			}
			if got.Fahrenheit() != tc.fahrenheit {
				t.Errorf("Expected %v, got %v", tc.fahrenheit, got.Fahrenheit())
			}
			if got.Gravity() != tc.gravity {
				t.Errorf("Expected %v, got %v", tc.gravity, got.Gravity())
			}
		})
	}
}
