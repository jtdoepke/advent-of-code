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
	assert.Equal(t, "5", got)
}

func Test_problemPart2(t *testing.T) {
	got := problemPart2(getTestData(t))
	assert.Equal(t, "12", got)
}

func Test_parseInput(t *testing.T) {
	got := parseInput(getTestData(t))
	want := []line{
		{
			p1: point{
				x: 0,
				y: 9,
			},
			p2: point{
				x: 5,
				y: 9,
			},
		},
		{
			p1: point{
				x: 0,
				y: 8,
			},
			p2: point{
				x: 8,
				y: 0,
			},
		},
		{
			p1: point{
				x: 3,
				y: 4,
			},
			p2: point{
				x: 9,
				y: 4,
			},
		},
		{
			p1: point{
				x: 2,
				y: 1,
			},
			p2: point{
				x: 2,
				y: 2,
			},
		},
		{
			p1: point{
				x: 7,
				y: 0,
			},
			p2: point{
				x: 7,
				y: 4,
			},
		},
		{
			p1: point{
				x: 2,
				y: 0,
			},
			p2: point{
				x: 6,
				y: 4,
			},
		},
		{
			p1: point{
				x: 0,
				y: 9,
			},
			p2: point{
				x: 2,
				y: 9,
			},
		},
		{
			p1: point{
				x: 1,
				y: 4,
			},
			p2: point{
				x: 3,
				y: 4,
			},
		},
		{
			p1: point{
				x: 0,
				y: 0,
			},
			p2: point{
				x: 8,
				y: 8,
			},
		},
		{
			p1: point{
				x: 5,
				y: 5,
			},
			p2: point{
				x: 8,
				y: 2,
			},
		},
	}
	assert.Equal(t, want, got)
}

func Test_part1Grid(t *testing.T) {
	lines := parseInput(getTestData(t))
	got := part1Grid(lines)
	want := grid{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	}
	assert.Equal(t, want, got)
}

func Test_part2Grid(t *testing.T) {
	lines := parseInput(getTestData(t))
	got := part2Grid(lines)
	want := grid{
		{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 0, 0, 2, 0, 0},
		{0, 0, 2, 0, 1, 0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 2, 0, 2, 0, 0},
		{0, 1, 1, 2, 3, 1, 3, 2, 1, 1},
		{0, 0, 0, 1, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	}
	assert.Equal(t, want, got)
}
