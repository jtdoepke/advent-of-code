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
	lines := ParseInput(inp)
	var count int
	for _, line := range lines {
		for _, digit := range line.Digits {
			switch digit.Bits() {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return strconv.Itoa(count)
}

func problemPart2(inp []byte) string {
	lines := ParseInput(inp)
	var sum int
	for _, line := range lines {
		m := sortSignals(line.Signals)
		sum += decode(m, line.Digits[3])
		sum += decode(m, line.Digits[2]) * 10
		sum += decode(m, line.Digits[1]) * 100
		sum += decode(m, line.Digits[0]) * 1000
	}
	return strconv.Itoa(sum)
}

type ProblemLine struct {
	Signals []SigSeg
	Digits  []SigSeg
}

func ParseInput(inp []byte) []ProblemLine {
	lines := bytes.Split(bytes.TrimSpace(inp), []byte("\n"))
	out := make([]ProblemLine, len(lines))
	for i, line := range lines {
		parts := bytes.Split(line, []byte(" "))
		seenPipe := false
		for _, part := range parts {
			switch {
			case part[0] == '|':
				seenPipe = true
			case !seenPipe:
				out[i].Signals = append(out[i].Signals, FromBytes(part))
			default:
				out[i].Digits = append(out[i].Digits, FromBytes(part))
			}
		}
	}
	return out
}

func sortSignals(signals []SigSeg) []SigSeg {
	signals = append([]SigSeg(nil), signals...)

	m := make([]SigSeg, 10)

	// First, find 1, 4, 7, and 8.
	i := 0
loop:
	for i < len(signals) {
		signal := signals[i]
		switch signal.Bits() {
		case 2:
			m[1] = signal
		case 3:
			m[7] = signal
		case 4:
			m[4] = signal
		case 7:
			m[8] = signal
		default:
			i++
			continue loop
		}
		signals = append(signals[:i], signals[i+1:]...)
	}

	// Find the rest.
	for _, signal := range signals {
		switch {
		case signal.Bits() == 6 && (signal&m[1]).Bits() == 2 && (signal&m[4]).Bits() == 3 && (signal&m[7]).Bits() == 3:
			m[0] = signal
		case signal.Bits() == 5 && (signal&m[1]).Bits() == 1 && (signal&m[4]).Bits() == 2 && (signal&m[7]).Bits() == 2:
			m[2] = signal
		case signal.Bits() == 5 && (signal&m[1]).Bits() == 2 && (signal&m[4]).Bits() == 3 && (signal&m[7]).Bits() == 3:
			m[3] = signal
		case signal.Bits() == 5 && (signal&m[1]).Bits() == 1 && (signal&m[4]).Bits() == 3 && (signal&m[7]).Bits() == 2:
			m[5] = signal
		case signal.Bits() == 6 && (signal&m[1]).Bits() == 1 && (signal&m[4]).Bits() == 3 && (signal&m[7]).Bits() == 2:
			m[6] = signal
		case signal.Bits() == 6 && (signal&m[1]).Bits() == 2 && (signal&m[4]).Bits() == 4 && (signal&m[7]).Bits() == 3:
			m[9] = signal
		}
	}

	return m
}

func decode(m []SigSeg, digit SigSeg) int {
	for i := 0; i < len(m); i++ {
		if m[i] == digit {
			return i
		}
	}
	log.Panic("unknown digit", digit)
	return 0
}
