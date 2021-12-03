package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestData(t *testing.T) []byte {
	f, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func Test_part1(t *testing.T) {
	got := part1(getTestData(t))
	assert.Equal(t, "198", got)
}

func Test_part2(t *testing.T) {
	got := part2(getTestData(t))
	assert.Equal(t, "230", got)
}

func Test_parseLinesToBinary(t *testing.T) {
	gotData, gotBitSize := parseLinesToBinary(getTestData(t))
	want := []uint64{
		4,  // 00100
		30, // 11110
		22, // 10110
		23, // 10111
		21, // 10101
		15, // 01111
		7,  // 00111
		28, // 11100
		16, // 10000
		25, // 11001
		2,  // 00010
		10, // 01010
	}
	assert.Equal(t, want, gotData)
	assert.Equal(t, 5, gotBitSize)
}

func Test_gammaRate(t *testing.T) {
	data, bitSize := parseLinesToBinary(getTestData(t))
	got := gammaRate(data, bitSize)
	assert.Equal(t, 22, got)
}

func Test_oxygenRating(t *testing.T) {
	data, bitSize := parseLinesToBinary(getTestData(t))
	got := oxygenRating(data, bitSize)
	assert.Equal(t, uint64(23), got)
}

func Test_co2ScrubberRating(t *testing.T) {
	data, bitSize := parseLinesToBinary(getTestData(t))
	got := co2ScrubberRating(data, bitSize)
	assert.Equal(t, uint64(10), got)
}
