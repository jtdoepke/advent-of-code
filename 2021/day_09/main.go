package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
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
	mat := parseInput(inp)
	lowPoints := findLowPoints(mat)
	var riskLevel int
	for _, p := range lowPoints {
		riskLevel += mat[p[0]][p[1]] + 1
	}
	return strconv.Itoa(riskLevel)
}

func problemPart2(inp []byte) string {
	mat := parseInput(inp)

	basinMap := make([][]int, len(mat))
	for i := range mat {
		basinMap[i] = make([]int, len(mat[i]))
	}

	id := 1
	for x, row := range mat {
		for y := range row {
			if row[y] != 9 && basinMap[x][y] == 0 {
				fill(mat, basinMap, x, y, id)
				id++
			}
		}
	}

	regionSizes := make(sort.IntSlice, id-1)
	for _, row := range basinMap {
		for _, v := range row {
			if v != 0 {
				regionSizes[v-1]++
			}
		}
	}
	sort.Sort(sort.Reverse(regionSizes))

	return strconv.Itoa(regionSizes[0] * regionSizes[1] * regionSizes[2])
}

func parseInput(inp []byte) (mat [][]int) {
	var row []int
	inp = bytes.TrimSpace(inp)
	buf := bytes.NewBuffer(make([]byte, 0, len(inp)))
	buf.Write(inp)
	for {
		b, err := buf.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panic(err)
		}
		switch b {
		case '0':
			row = append(row, 0)
		case '1':
			row = append(row, 1)
		case '2':
			row = append(row, 2)
		case '3':
			row = append(row, 3)
		case '4':
			row = append(row, 4)
		case '5':
			row = append(row, 5)
		case '6':
			row = append(row, 6)
		case '7':
			row = append(row, 7)
		case '8':
			row = append(row, 8)
		case '9':
			row = append(row, 9)
		case '\n':
			mat = append(mat, row)
			row = nil
		}
	}
	if len(row) > 0 {
		mat = append(mat, row)
	}
	return
}

type Pair [2]int

func findLowPoints(mat [][]int) (lowPoints []Pair) {
	for x, row := range mat {
		for y, v := range row {
			lowPoint := true

			if x > 0 && mat[x-1][y] <= v {
				lowPoint = false
			}
			if x < len(mat)-1 && mat[x+1][y] <= v {
				lowPoint = false
			}
			if y > 0 && mat[x][y-1] <= v {
				lowPoint = false
			}
			if y < len(row)-1 && mat[x][y+1] <= v {
				lowPoint = false
			}

			if lowPoint {
				lowPoints = append(lowPoints, Pair{x, y})
			}
		}
	}
	return
}

func fill(mat, basinMap [][]int, x, y, v int) {
	shouldFill := func(x, y int) bool {
		if x < 0 || x >= len(mat) || y < 0 || y >= len(mat[x]) {
			return false
		}
		return mat[x][y] != 9 && basinMap[x][y] == 0
	}
	if shouldFill(x, y) {
		basinMap[x][y] = v
	}
	if shouldFill(x-1, y) {
		fill(mat, basinMap, x-1, y, v)
	}
	if shouldFill(x, y+1) {
		fill(mat, basinMap, x, y+1, v)
	}
	if shouldFill(x+1, y) {
		fill(mat, basinMap, x+1, y, v)
	}
	if shouldFill(x, y-1) {
		fill(mat, basinMap, x, y-1, v)
	}
}
