package tilt

import (
	"context"
	"log"
	"time"

	"github.com/JuulLabs-OSS/ble"
	"github.com/JuulLabs-OSS/ble/examples/lib/dev"
	"github.com/pkg/errors"
)

// Scanner for Tilt devices
type Scanner struct {
	devices Devices
        d ble.Device
}

// Devices stores discovered devices
type Devices map[Colour]Tilt

// NewScanner returns a Scanner
func NewScanner() *Scanner {
	return &Scanner{}
}

// Scan finds Tilt devices and times out after a duration
func (s *Scanner) Scan(timeout time.Duration) {

	log.Printf("Scanning for %v", timeout)

	s.devices = make(map[Colour]Tilt)
        var err error = nil

	if s.d == nil {
	    s.d, err = dev.NewDevice("go-tilt")
	    if err != nil {
		log.Fatalf("Unable to initialise new device : %s", err)
	    }
	    ble.SetDefaultDevice(s.d)
        }
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), timeout))
	err = ble.Scan(ctx, false, s.advHandler, advFilter)
	if err != nil {
		switch errors.Cause(err) {
		case nil:
		case context.DeadlineExceeded:
			log.Printf("Finished scanning\n")
		case context.Canceled:
			log.Printf("Cancelled\n")
		default:
			log.Fatalf(err.Error())
		}
	}
}

func advFilter(a ble.Advertisement) bool {
	return IsTilt(a.ManufacturerData())
}

func (s *Scanner) advHandler(a ble.Advertisement) {

	// create iBeacon
	b, err := NewIBeacon(a.ManufacturerData())
	if err != nil {
		log.Println(err)
		return
	}

	// create Tilt from iBeacon
	t, err := NewTilt(b)
	if err != nil {
		log.Println(err)
		return
	}

	s.HandleTilt(t)
}

// HandleTilt adds a discovered Tilt to a map
func (s *Scanner) HandleTilt(t Tilt) {
	s.devices[t.col] = t
}

// Tilts contains the found devices
func (s *Scanner) Tilts() Devices {
	return s.devices
}
