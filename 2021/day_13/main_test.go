package main

import (
	"io"
	"os"
	"strings"
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
	assert.Equal(t, "17", got)
}

func Test_problemPart2(t *testing.T) {
	got := problemPart2(getTestData(t))
	want := strings.TrimSpace(`
#####
#...#
#...#
#...#
#####
.....
.....
`)
	assert.Equal(t, want, got)
}

func TestSheet_String(t *testing.T) {
	sheet, _ := parseInput(getTestData(t))
	got := sheet.String()
	want := strings.TrimSpace(`
...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
...........
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
`)
	assert.Equal(t, want, got)
}

func TestSheet_Fold(t *testing.T) {
	sheet, folds := parseInput(getTestData(t))
	got1 := sheet.Fold(folds[0])
	want1 := strings.TrimSpace(`
#.##..#..#.
#...#......
......#...#
#...#......
.#.#..#.###
...........
...........
`)
	assert.Equal(t, want1, got1.String())

	got2 := got1.Fold(folds[1])
	want2 := strings.TrimSpace(`
#####
#...#
#...#
#...#
#####
.....
.....
`)
	assert.Equal(t, want2, got2.String())
}
