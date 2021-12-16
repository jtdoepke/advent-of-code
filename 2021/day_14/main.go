package main

import (
	"bytes"
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
	counts, pairs, pairInsertionRules := parseInput(inp)
	counts, pairs = doPairInsertions(10, counts, pairs, pairInsertionRules)

	var (
		mostCommonCount  int
		leastCommonCount = math.MaxInt
	)
	for _, count := range counts {
		if count > mostCommonCount {
			mostCommonCount = count
		}
		if count < leastCommonCount {
			leastCommonCount = count
		}
	}

	return strconv.Itoa(mostCommonCount - leastCommonCount)
}

func problemPart2(inp []byte) string {
	counts, pairs, pairInsertionRules := parseInput(inp)
	counts, pairs = doPairInsertions(40, counts, pairs, pairInsertionRules)

	var (
		mostCommonCount  int
		leastCommonCount = math.MaxInt
	)
	for _, count := range counts {
		if count > mostCommonCount {
			mostCommonCount = count
		}
		if count < leastCommonCount {
			leastCommonCount = count
		}
	}

	return strconv.Itoa(mostCommonCount - leastCommonCount)
}

func doPairInsertions(n int, counts map[byte]int, pairs map[string]int, pairInsertionRules map[string]byte) (map[byte]int, map[string]int) {
	for step := 0; step < n; step++ {
		newPairs := make(map[string]int, len(pairs))
		for pair, count := range pairs {
			if insert, ok := pairInsertionRules[pair]; ok {
				counts[insert] += count
				newPairs[string(pair[0])+string(insert)] += count
				newPairs[string(insert)+string(pair[1])] += count
			} else {
				newPairs[pair] += count
			}
		}
		pairs = newPairs
	}

	var (
		mostCommonCount  int
		leastCommonCount = math.MaxInt
	)
	for _, count := range counts {
		if count > mostCommonCount {
			mostCommonCount = count
		}
		if count < leastCommonCount {
			leastCommonCount = count
		}
	}

	return counts, pairs
}

func parseInput(inp []byte) (elementCounts map[byte]int, pairs map[string]int, pairInsertionRules map[string]byte) {
	lines := bytes.Split(bytes.TrimSpace(inp), []byte("\n"))
	elementCounts = make(map[byte]int)
	pairs = make(map[string]int)
	for i := 0; i < len(lines[0])-1; i++ {
		elementCounts[lines[0][i]]++
		pairs[string(lines[0][i:i+2])]++
	}
	if len(lines[0]) > 1 {
		elementCounts[lines[0][len(lines[0])-1]]++
	}
	pairInsertionRules = make(map[string]byte, len(lines)-2)
	for i := 2; i < len(lines); i++ {
		parts := bytes.Split(lines[i], []byte(" -> "))
		pairInsertionRules[string(parts[0])] = parts[1][0]
	}
	return
}
