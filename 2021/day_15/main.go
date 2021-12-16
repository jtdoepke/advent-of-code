package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	astar "github.com/beefsack/go-astar"
)

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
	grid := parseInput(inp)
	lastRow := grid[len(grid)-1]
	path, distance, found := astar.Path(&grid[0][0], &lastRow[len(lastRow)-1])
	if !found {
		log.Panic("not found")
	}
	keep(path)
	// for _, p := range path {
	// 	t := p.(*Tile)
	// 	fmt.Printf("%d,%d\n", t.x, t.y)
	// }
	return strconv.Itoa(int(distance))
}

func problemPart2(inp []byte) string {
	grid := parseInput(inp)
	grid = repeat(grid, 5, 5)
	// fmt.Println(grid.String())
	lastRow := grid[len(grid)-1]
	path, distance, found := astar.Path(&grid[0][0], &lastRow[len(lastRow)-1])
	if !found {
		log.Panic("not found")
	}
	keep(path)
	// for _, p := range path {
	// 	t := p.(*Tile)
	// 	fmt.Printf("%d,%d\n", t.x, t.y)
	// }
	return strconv.Itoa(int(distance))
}

type Tile struct {
	grid [][]Tile
	x, y int
	risk int
}

func (t *Tile) PathNeighbors() (neighbors []astar.Pather) {
	if t.x > 0 {
		neighbors = append(neighbors, &t.grid[t.x-1][t.y])
	}
	if t.x < len(t.grid)-1 {
		neighbors = append(neighbors, &t.grid[t.x+1][t.y])
	}
	if t.y > 0 {
		neighbors = append(neighbors, &t.grid[t.x][t.y-1])
	}
	if t.y < len(t.grid[t.x])-1 {
		neighbors = append(neighbors, &t.grid[t.x][t.y+1])
	}
	return
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	toTile := to.(*Tile)
	return float64(toTile.risk)
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toTile := to.(*Tile)
	return float64(manhattenDistance(t.x, toTile.x, t.y, toTile.y))
}

type Grid [][]Tile

func (g Grid) String() string {
	buf := bytes.NewBuffer(make([]byte, len(g)*len(g[0])+len(g)))
	for _, row := range g {
		for _, t := range row {
			switch t.risk {
			case 0:
				buf.WriteByte('0')
			case 1:
				buf.WriteByte('1')
			case 2:
				buf.WriteByte('2')
			case 3:
				buf.WriteByte('3')
			case 4:
				buf.WriteByte('4')
			case 5:
				buf.WriteByte('5')
			case 6:
				buf.WriteByte('6')
			case 7:
				buf.WriteByte('7')
			case 9:
				buf.WriteByte('9')
			case 8:
				buf.WriteByte('8')
			}
		}
		buf.WriteByte('\n')
	}
	return string(bytes.Trim(buf.Bytes(), "\n\x00"))
}

func parseInput(inp []byte) (grid Grid) {
	lines := bytes.Split(bytes.TrimSpace(inp), []byte("\n"))
	grid = make([][]Tile, len(lines))
	for x, line := range lines {
		grid[x] = make([]Tile, len(line))
		for y, c := range line {
			risk, err := strconv.Atoi(string(c))
			if err != nil {
				log.Panic(err)
			}
			grid[x][y] = Tile{
				grid: grid,
				x:    x,
				y:    y,
				risk: risk,
			}
		}
	}
	return
}

func keep(x interface{}) {
	// I want to keep the debug logging commented out,
	// but I don't want the Go compiler complaining about
	// unused vars.
}

func repeat(grid Grid, timesX, timesY int) Grid {
	origX, origY := len(grid), len(grid[0])
	newGrid := make(Grid, origX*timesX)
	for i := range newGrid {
		newGrid[i] = make([]Tile, origY*timesY)
	}
	for x := range newGrid {
		for y := range newGrid[x] {
			risk := grid[x%origX][y%origY].risk
			incBy := manhattenDistance(0, 0, x/origX, y/origY)
			risk = (((risk - 1) + incBy) % 9) + 1
			newGrid[x][y] = Tile{
				grid: newGrid,
				x:    x,
				y:    y,
				risk: risk,
			}
		}
	}
	return newGrid
}

func manhattenDistance(x1, y1, x2, y2 int) int {
	absX := x1 - x2
	if absX < 0 {
		absX = -absX
	}
	absY := y1 - y2
	if absY < 0 {
		absY = -absY
	}
	return absX + absY
}
