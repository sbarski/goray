package main

import (
	"testing"
)

func TestPPMImageRead(t *testing.T) {
	read := "west.ppm"
	write := "west_test.ppm"

	pixels, width, height, depth, err := ReadPPM(read)

	if err != nil {
		t.Fatalf("Could not read file: ", err)
	}

	WriteToPPM(pixels, width, height, depth, write)
}
