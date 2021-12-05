package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var inputPattern = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	b = bytes.TrimSpace(b)
	fmt.Println("Part 1 Answer:")
	fmt.Println(problemPart1(b))
	fmt.Println("-------------------------------")
	fmt.Println("Part 2 Answer:")
	fmt.Println(problemPart2(b))
}

func problemPart1(inp []byte) string {
	lines := parseInput(inp)
	g := part1Grid(lines)
	var dangerZones int
	for _, row := range g {
		for _, cel := range row {
			if cel > 1 {
				dangerZones++
			}
		}
	}
	return strconv.Itoa(dangerZones)
}

func part1Grid(lines []line) grid {
	g := newGridFromLines(lines)
	for _, line := range lines {
		if line.xEqual() {
			for i := line.p1.y; i <= line.p2.y; i++ {
				g[i][line.p1.x]++
			}
		} else if line.yEqual() {
			for i := line.p1.x; i <= line.p2.x; i++ {
				g[line.p1.y][i]++
			}
		}
	}
	return g
}

func problemPart2(inp []byte) string {
	lines := parseInput(inp)
	g := part2Grid(lines)
	var dangerZones int
	for _, row := range g {
		for _, cel := range row {
			if cel > 1 {
				dangerZones++
			}
		}
	}
	return strconv.Itoa(dangerZones)
}

func part2Grid(lines []line) grid {
	g := newGridFromLines(lines)
	for _, line := range lines {
		if line.xEqual() {
			for i := line.p1.y; i <= line.p2.y; i++ {
				g[i][line.p1.x]++
			}
		} else if line.yEqual() {
			for i := line.p1.x; i <= line.p2.x; i++ {
				g[line.p1.y][i]++
			}
		} else {
			// Diagonal line
			var xOp, yOp int
			if line.p1.x < line.p2.x {
				xOp = 1
			} else {
				xOp = -1
			}
			if line.p1.y < line.p2.y {
				yOp = 1
			} else {
				yOp = -1
			}
			x, y := line.p1.x, line.p1.y
			for {
				g[y][x]++
				if x == line.p2.x && y == line.p2.y {
					break
				}
				x, y = x+xOp, y+yOp
			}
		}
	}
	return g
}

type point struct {
	x int
	y int
}

type line struct {
	p1 point
	p2 point
}

func (ln line) xEqual() bool {
	return ln.p1.x == ln.p2.x
}

func (ln line) yEqual() bool {
	return ln.p1.y == ln.p2.y
}

func parseInput(inp []byte) []line {
	matches := inputPattern.FindAllStringSubmatch(string(inp), -1)
	lines := make([]line, len(matches))
	for i, match := range matches {
		lines[i].p1.x = mustAtoi(match[1])
		lines[i].p1.y = mustAtoi(match[2])
		lines[i].p2.x = mustAtoi(match[3])
		lines[i].p2.y = mustAtoi(match[4])
		if lines[i].p1.x > lines[i].p2.x || (lines[i].p1.x == lines[i].p2.x && lines[i].p1.y > lines[i].p2.y) {
			lines[i].p2, lines[i].p1 = lines[i].p1, lines[i].p2
		}
	}
	return lines
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type grid [][]int

func newGrid(sizeX, sizeY int) grid {
	g := make(grid, sizeY)
	for i := 0; i < sizeY; i++ {
		g[i] = make([]int, sizeX)
	}
	return g
}

func newGridFromLines(lines []line) grid {
	var max int
	for _, line := range lines {
		if line.p1.x > max {
			max = line.p1.x
		}
		if line.p2.x > max {
			max = line.p2.x
		}
		if line.p1.y > max {
			max = line.p1.y
		}
		if line.p2.y > max {
			max = line.p2.x
		}
	}
	return newGrid(max+1, max+1)
}
