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

var (
	syntaxPoints = map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	completionPoints = map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
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
	lines := bytes.Split(bytes.TrimSpace(inp), []byte("\n"))
	var score int
loop:
	for _, line := range lines {
		stack := make(CharStack, 0)
		for _, b := range line {
			switch b {
			case '[':
				stack.Push(']')
			case '(':
				stack.Push(')')
			case '<':
				stack.Push('>')
			case '{':
				stack.Push('}')
			default:
				expected := stack.Pop()
				if b != expected {
					score += syntaxPoints[b]
					continue loop
				}
			}
		}
	}
	return strconv.Itoa(score)
}

func problemPart2(inp []byte) string {
	lines := bytes.Split(bytes.TrimSpace(inp), []byte("\n"))
	var scores []int

loop:
	for _, line := range lines {
		stack := make(CharStack, 0)
		for _, b := range line {
			switch b {
			case '[':
				stack.Push(']')
			case '(':
				stack.Push(')')
			case '<':
				stack.Push('>')
			case '{':
				stack.Push('}')
			default:
				expected := stack.Pop()
				if b != expected {
					// Corrupted, skip it.
					continue loop
				}
			}
		}
		var lineScore int
		for i := len(stack) - 1; i >= 0; i-- {
			lineScore *= 5
			lineScore += completionPoints[stack[i]]
		}
		scores = append(scores, lineScore)
	}

	sort.Ints(scores)

	return strconv.Itoa(scores[len(scores)/2])
}

// type Chunk struct {
// 	OpenChar  byte
// 	Subchunks []*Chunk
// }

type CharStack []byte

func (cs *CharStack) Push(b byte) {
	*cs = append(*cs, b)
}
func (cs *CharStack) Pop() byte {
	s := *cs
	if len(s) == 0 {
		return 0
	}
	b := s[len(s)-1]
	*cs = s[:len(s)-1]
	return b
}
