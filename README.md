# Tilt API in Go

[![GoDoc](https://godoc.org/github.com/alexhowarth/go-tilt?status.svg)](https://pkg.go.dev/github.com/alexhowarth/go-tilt)
[![Go Report Card](https://goreportcard.com/badge/alexhowarth/go-tilt)](https://goreportcard.com/report/github.com/alexhowarth/go-tilt)

This package provides a library for reading data from Tilt sensors.

Tilt sensors are devices for brewers that read specific gravity and temperature during fermentation.

## Installation

~~~~
go get github.com/alexhowarth/go-tilt
~~~~

## Usage

See the examples directory.

## Compatibility

This project currently uses a specific fork of go-ble that is compatible with both Linux and macOS.

## Cross compiling

To create a binary for use on a Raspberry Pi, simply build it for ARM and copy the binary (no other dependencies are required).

~~~~
env GOOS=linux GOARCH=arm go build examples/scanner/scanner.go
~~~~
