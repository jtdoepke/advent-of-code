package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
	sheet, folds := parseInput(inp)
	sheet = sheet.Fold(folds[0])
	var count int
	for _, row := range sheet {
		for _, v := range row {
			if v {
				count++
			}
		}
	}
	// log.Println(sheet.String())
	return strconv.Itoa(count)
}

func problemPart2(inp []byte) string {
	sheet, folds := parseInput(inp)
	for _, f := range folds {
		sheet = sheet.Fold(f)
	}
	var count int
	for _, row := range sheet {
		for _, v := range row {
			if v {
				count++
			}
		}
	}
	return sheet.String()
}

type Sheet [][]bool

func (s Sheet) String() string {
	if len(s) == 0 {
		return ""
	}
	lines := make([][]byte, len(s[0]))
	for i := range s[0] {
		lines[i] = make([]byte, len(s))
	}
	for x := range s {
		for y, v := range s[x] {
			if v {
				lines[y][x] = '#'
			} else {
				lines[y][x] = '.'
			}
		}
	}
	b := strings.Builder{}
	for _, line := range lines {
		b.Write(line)
		b.WriteByte('\n')
	}
	return strings.TrimSpace(b.String())
}

func (s Sheet) Fold(f Fold) Sheet {
	var newSheet Sheet
	if f.Axis == 'x' {
		newXSize := max(f.Pos, len(s)-f.Pos-1)
		newSheet = make(Sheet, newXSize)
		for i := 0; i < f.Pos; i++ {
			newSheet[i] = append([]bool(nil), s[i]...)
		}
		for i := f.Pos + 1; i < len(s); i++ {
			for j := range s[i] {
				if s[i][j] {
					newSheet[f.Pos-(i-f.Pos)][j] = true
				}
			}
		}
	} else if f.Axis == 'y' {
		newYSize := max(f.Pos, len(s[0])-f.Pos-1)
		newSheet = make(Sheet, len(s))
		for i, col := range s {
			newSheet[i] = make([]bool, newYSize)
			copy(newSheet[i], col[:f.Pos])
			for j := f.Pos + 1; j < len(col); j++ {
				if col[j] {
					newSheet[i][f.Pos-(j-f.Pos)] = true
				}
			}
		}
	} else {
		log.Panic("bad axis", f.Axis)
	}
	return newSheet
}

type Fold struct {
	Axis byte
	Pos  int
}

func parseInput(inp []byte) (sheet Sheet, folds []Fold) {
	var maxX, maxY int
	lines := bytes.Split(bytes.TrimSpace(inp), []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		switch line[0] {
		case 'f':
			f := Fold{
				Axis: line[11],
			}
			pos, err := strconv.Atoi(string(line[13:]))
			if err != nil {
				log.Panic(err)
			}
			f.Pos = pos
			folds = append(folds, f)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			parts := bytes.Split(line, []byte(","))

			x, err := strconv.Atoi(string(parts[0]))
			if err != nil {
				log.Panic(err)

			}
			y, err := strconv.Atoi(string(parts[1]))
			if err != nil {
				log.Panic(err)
			}

			if maxX < x {
				maxX = x
			}
			if maxY < y {
				maxY = y
			}

			for len(sheet) < maxX+1 {
				sheet = append(sheet, make([]bool, maxY))
			}
			if len(sheet[x]) < maxY+1 {
				newY := make([]bool, maxY+1)
				copy(newY, sheet[x])
				sheet[x] = newY
			}

			sheet[x][y] = true
		}
	}
	for i := range sheet {
		if len(sheet[i]) < maxY+1 {
			newY := make([]bool, maxY+1)
			copy(newY, sheet[i])
			sheet[i] = newY
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
