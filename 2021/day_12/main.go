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
	caves := parseInput(inp)
	paths := dfsNoSmall(caves, Cave("start"), nil)
	// for _, p := range paths {
	// 	log.Println(p.String())
	// }
	return strconv.Itoa(len(paths))
}

func problemPart2(inp []byte) string {
	caves := parseInput(inp)
	paths := dfsOneSmall(caves, Cave("start"), nil)
	// for _, p := range paths {
	// 	log.Println(p.String())
	// }
	return strconv.Itoa(len(paths))
}

func dfsNoSmall(caves map[Cave]CaveSet, start Cave, visited Path) (paths []Path) {
	if visited == nil {
		visited = Path{start}
	}
	for neighbor := range caves[start] {
		if neighbor.IsStart() {
			continue
		}
		if neighbor.IsEnd() {
			paths = append(paths, append(visited.Clone(), neighbor))
			continue
		}
		if neighbor.IsBig() || !visited.Contains(neighbor) {
			paths = append(paths, dfsNoSmall(
				caves,
				neighbor,
				append(visited.Clone(), neighbor),
			)...)
		}
	}
	return paths
}

func dfsOneSmall(caves map[Cave]CaveSet, start Cave, visited Path) (paths []Path) {
	if visited == nil {
		visited = Path{start}
	}
	for neighbor := range caves[start] {
		if neighbor.IsStart() {
			continue
		}
		if neighbor.IsEnd() {
			paths = append(paths, append(visited.Clone(), neighbor))
			continue
		}
		if neighbor.IsBig() || !visited.HasMoreThanOneSmall() || !visited.Contains(neighbor) {
			paths = append(paths, dfsOneSmall(
				caves,
				neighbor,
				append(visited.Clone(), neighbor),
			)...)
		}
	}
	return paths
}

type Cave string

func (c Cave) IsBig() bool {
	return strings.ToUpper(string(c)) == string(c)
}

func (c Cave) IsEnd() bool {
	return strings.EqualFold(string(c), "end")
}

func (c Cave) IsStart() bool {
	return strings.EqualFold(string(c), "start")
}

type CaveSet map[Cave]struct{}

func (cs CaveSet) Push(c Cave) {
	cs[c] = struct{}{}
}

func (cs CaveSet) Contains(c Cave) bool {
	_, ok := cs[c]
	return ok
}

type Path []Cave

func (p Path) Contains(c Cave) bool {
	for i := range p {
		if c == p[i] {
			return true
		}
	}
	return false
}

func (p Path) HasMoreThanOneSmall() bool {
	counts := make(map[Cave]int, len(p))
	for _, c := range p {
		if c.IsBig() || c.IsStart() || c.IsEnd() {
			continue
		}
		counts[c]++
	}
	for _, c := range counts {
		if c > 1 {
			return true
		}
	}
	return false
}

func (p Path) Clone() Path {
	p2 := make(Path, len(p))
	copy(p2, p)
	return p2
}

func (p Path) String() string {
	asStrings := make([]string, len(p))
	for i, c := range p {
		asStrings[i] = string(c)
	}
	return strings.Join(asStrings, ",")
}

func parseInput(inp []byte) map[Cave]CaveSet {
	lines := bytes.Split(bytes.TrimSpace(inp), []byte("\n"))
	caves := make(map[Cave]CaveSet, len(lines)*2)
	for _, line := range lines {
		parts := bytes.Split(line, []byte("-"))

		c1 := Cave(parts[0])
		if _, ok := caves[c1]; !ok {
			caves[c1] = make(CaveSet)
		}

		c2 := Cave(parts[1])
		if _, ok := caves[c2]; !ok {
			caves[c2] = make(CaveSet)
		}

		caves[c1].Push(c2)
		caves[c2].Push(c1)
	}
	return caves
}
