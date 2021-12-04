package main

import (
	"bufio"
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
	fmt.Println("Part 1 Answer:")
	fmt.Println(problemPart1(b))
	fmt.Println("-------------------------------")
	fmt.Println("Part 2 Answer:")
	fmt.Println(problemPart2(b))
}

func problemPart1(inp []byte) string {
	calls, boards, err := parseInput(inp)
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range calls {
		for _, board := range boards {
			board.Call(n)
			if board.HasWon() {
				return strconv.Itoa(board.Score(n))
			}
		}
	}
	panic("no boards won")
}

func problemPart2(inp []byte) string {
	calls, boards, err := parseInput(inp)
	if err != nil {
		log.Fatal(err)
	}

	hasWon := make([]bool, len(boards))

	for _, n := range calls {
		for i, board := range boards {
			if !hasWon[i] {
				board.Call(n)
				if board.HasWon() {
					hasWon[i] = true
					if allTrue(hasWon) {
						return strconv.Itoa(board.Score(n))
					}
				}
			}
		}
	}
	panic("no boards won")
}

func parseInput(inp []byte) (calls []int, boards []*BingoBoard, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(inp))
	var currentBoard [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if calls == nil { // Parse the first line.
			nums := strings.Split(line, ",")
			for _, n := range nums {
				ni, err := strconv.Atoi(n)
				if err != nil {
					return nil, nil, err
				}
				calls = append(calls, ni)
			}

		} else if line == "" {
			if currentBoard != nil {
				boards = append(boards, NewBingoBoard(currentBoard))
				currentBoard = nil
			}
		} else {
			var nums []int
			for _, n := range strings.Fields(line) {
				ni, err := strconv.Atoi(n)
				if err != nil {
					return nil, nil, err
				}
				nums = append(nums, ni)
			}
			currentBoard = append(currentBoard, nums)
		}
	}
	if currentBoard != nil {
		boards = append(boards, NewBingoBoard(currentBoard))
		currentBoard = nil
	}
	return
}

func allTrue(bs []bool) bool {
	for i := range bs {
		if !bs[i] {
			return false
		}
	}
	return true
}
