package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
	squids := parseInput(inp)
	toFlash := make(coordStack)
	toZero := make(coordStack)
	var totalFlashes int
	for step := 0; step < 100; step++ {
		// Phase 1
		for x, row := range squids {
			for y := range row {
				row[y]++
				if row[y] > 9 {
					toFlash.Push([2]int{x, y})
				}
			}
		}

		// Phase 2
		for len(toFlash) > 0 {
			c := toFlash.Pop()
			if _, ok := toZero[c]; ok {
				continue
			}
			toZero.Push(c)
			x, y := c[0], c[1]
			for x1 := x - 1; x1 <= x+1; x1++ {
				for y1 := y - 1; y1 <= y+1; y1++ {
					if x1 == x && y1 == y {
						continue
					}
					if x1 < 0 || x1 >= len(squids) {
						continue
					}
					if y1 < 0 || y1 >= len(squids[x1]) {
						continue
					}
					squids[x1][y1]++
					if squids[x1][y1] > 9 {
						toFlash.Push([2]int{x1, y1})
					}
				}
			}
		}

		// Phase 3
		totalFlashes += len(toZero)
		for len(toZero) > 0 {
			c := toZero.Pop()
			squids[c[0]][c[1]] = 0
		}
	}

	return strconv.Itoa(totalFlashes)
}

func problemPart2(inp []byte) string {
	squids := parseInput(inp)
	stopAt := size(squids)
	toFlash := make(coordStack)
	toZero := make(coordStack)
	for step := 0; step < 1000000; step++ {
		// Phase 1
		for x, row := range squids {
			for y := range row {
				row[y]++
				if row[y] > 9 {
					toFlash.Push([2]int{x, y})
				}
			}
		}

		// Phase 2
		for len(toFlash) > 0 {
			c := toFlash.Pop()
			if _, ok := toZero[c]; ok {
				continue
			}
			toZero.Push(c)
			x, y := c[0], c[1]
			for x1 := x - 1; x1 <= x+1; x1++ {
				for y1 := y - 1; y1 <= y+1; y1++ {
					if x1 == x && y1 == y {
						continue
					}
					if x1 < 0 || x1 >= len(squids) {
						continue
					}
					if y1 < 0 || y1 >= len(squids[x1]) {
						continue
					}
					squids[x1][y1]++
					if squids[x1][y1] > 9 {
						toFlash.Push([2]int{x1, y1})
					}
				}
			}
		}

		// Phase 3
		if len(toZero) == stopAt {
			return strconv.Itoa(step + 1)
		}
		for len(toZero) > 0 {
			c := toZero.Pop()
			squids[c[0]][c[1]] = 0
		}
	}

	return ""
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

type coordStack map[[2]int]struct{}

func (s *coordStack) Push(coords [2]int) {
	(*s)[coords] = struct{}{}
}

func (s *coordStack) Pop() [2]int {
	for k := range *s {
		delete(*s, k)
		return k
	}
	return [2]int{}
}

func size(mat [][]int) int {
	var size int
	for _, row := range mat {
		size += len(row)
	}
	return size
}
