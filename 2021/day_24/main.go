package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
	return "NOT IMPLEMENTED"
}

func problemPart2(inp []byte) string {
	return "NOT IMPLEMENTED"
}
