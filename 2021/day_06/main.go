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
	fishCounts := parseInput(inp)
	result := simulateFish(fishCounts, 80)
	return strconv.Itoa(result)
}

func problemPart2(inp []byte) string {
	fishCounts := parseInput(inp)
	result := simulateFish(fishCounts, 256)
	return strconv.Itoa(result)
}

func simulateFish(fishCounts []int, days int) int {
	const (
		spawnCountdown = 8
		resetCountdown = 6
	)

	for len(fishCounts) < spawnCountdown+1 {
		fishCounts = append(fishCounts, 0)
	}

	// log.Printf("Initial state: %d", sum(fishCounts))
	for day := 0; day < days; day++ {
		spawningFish := fishCounts[0]
		for i := 1; i < len(fishCounts); i++ {
			fishCounts[i-1] = fishCounts[i]
		}
		fishCounts[len(fishCounts)-1] = 0
		fishCounts[resetCountdown] += spawningFish
		fishCounts[spawnCountdown] += spawningFish
		// log.Printf("After %d days: %d", day+1, sum(fishCounts))
	}
	return sum(fishCounts)
}

func sum(its []int) int {
	var sum int
	for _, c := range its {
		sum += c
	}
	return sum
}

func parseInput(b []byte) (fishCounts []int) {
	parts := strings.Split(string(b), ",")
	for _, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err)
		}
		for len(fishCounts) < (n + 1) {
			fishCounts = append(fishCounts, 0)
		}
		fishCounts[n]++
	}
	return
}
