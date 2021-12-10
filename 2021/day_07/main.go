package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
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
	positions := parseInput(inp)
	var max int
	for _, pos := range positions {
		if pos > max {
			max = pos
		}
	}
	costs := make([]int, max+1)
	for _, pos := range positions {
		for i := 0; i < max+1; i++ {
			costs[i] += abs(pos - i)
		}
	}
	var minCost int
	minCost = math.MaxInt
	for _, cost := range costs {
		if cost < minCost {
			minCost = cost
		}
	}
	return strconv.Itoa(minCost)
}

func problemPart2(inp []byte) string {
	positions := parseInput(inp)
	var max int
	for _, pos := range positions {
		if pos > max {
			max = pos
		}
	}
	costs := make([]int, max+1)
	for _, pos := range positions {
		for i := 0; i < max+1; i++ {
			costs[i] += gaussSum(abs(pos - i))
		}
	}
	var minCost int
	minCost = math.MaxInt
	for _, cost := range costs {
		if cost < minCost {
			minCost = cost
		}
	}
	return strconv.Itoa(minCost)
}

func parseInput(inp []byte) (positions []int) {
	inp = bytes.TrimSpace(inp)
	b := bytes.NewBuffer(make([]byte, 0, len(inp)+2))
	b.WriteByte('[')
	b.Write(inp)
	b.WriteByte(']')
	if err := json.NewDecoder(b).Decode(&positions); err != nil {
		log.Panic(err)
	}
	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func gaussSum(n int) int {
	return (n * (n + 1)) / 2
}
