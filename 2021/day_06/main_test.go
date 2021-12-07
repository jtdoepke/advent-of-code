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

func Test_problemPart1(t *testing.T) {
	got := problemPart1(getTestData(t))
	assert.Equal(t, "5934", got)
}

func Test_problemPart2(t *testing.T) {
	got := problemPart2(getTestData(t))
	assert.Equal(t, "NOT IMPLEMENTED", got)
}
