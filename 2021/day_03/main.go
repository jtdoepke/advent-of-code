package main

import (
	"bufio"
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
	fmt.Println(part1(b))
	fmt.Println("-------------------------------")
	fmt.Println(part2(b))
}

func part1(inp []byte) string {
	data, bitSize := parseLinesToBinary(inp)
	gamma := gammaRate(data, bitSize)
	epsilon := gamma ^ makeMask(bitSize)
	return strconv.Itoa(gamma * epsilon)
}

func part2(inp []byte) string {
	data, bitSize := parseLinesToBinary(inp)
	oR := oxygenRating(data, bitSize)
	co2R := co2ScrubberRating(data, bitSize)
	result := int(oR) * int(co2R)
	return strconv.Itoa(result)
}

func parseLinesToBinary(b []byte) (data []uint64, bitSize int) {
	sc := bufio.NewScanner(bytes.NewReader(b))
	for sc.Scan() {
		line := sc.Text()
		bitSize = len(line)
		i, err := strconv.ParseInt(line, 2, 64)
		if err != nil {
			panic(err)
		}
		data = append(data, uint64(i))
	}
	return data, bitSize
}

func gammaRate(data []uint64, bitSize int) (gamma int) {
	for idx := 0; idx < bitSize; idx++ {
		ones := countTrue(bitColumn(data, idx))
		if ones >= len(data)-ones {
			gamma |= 1 << idx
		}
	}
	return
}

func bitColumn(data []uint64, column int) []bool {
	out := make([]bool, len(data))
	for idx, p := range data {
		b := p & (1 << column)
		out[idx] = (b != 0)
	}
	return out
}

func countTrue(bools []bool) (i int) {
	for _, b := range bools {
		if b {
			i += 1
		}
	}
	return
}

func makeMask(size int) (mask int) {
	for i := 0; i < size; i++ {
		mask |= 1 << i
	}
	return
}

type FilterCriteria func(count0, count1 int) bool

func filterData(data []uint64, bitSize int, criteria FilterCriteria) uint64 {
	data = append([]uint64(nil), data...) // Make a copy.
	for column := bitSize - 1; column >= 0; column-- {
		c := bitColumn(data, column)
		ones := countTrue(c)
		keep := criteria(len(data)-ones, ones)

		n := 0
		for i := 0; i < len(data); i++ {
			if c[i] == keep {
				data[n] = data[i]
				n++
			}
		}
		data = data[:n]

		if len(data) == 1 {
			return data[0]
		} else if len(data) == 0 {
			panic("data completely filtered out")
		}
	}
	panic("checked all columns")
}

func oxygenRating(data []uint64, bitSize int) uint64 {
	return filterData(data, bitSize, func(count0, count1 int) bool {
		if count1 > count0 {
			return true
		} else if count0 > count1 {
			return false
		}
		return true
	})
}

func co2ScrubberRating(data []uint64, bitSize int) uint64 {
	return filterData(data, bitSize, func(count0, count1 int) bool {
		if count1 > count0 {
			return false
		} else if count0 > count1 {
			return true
		}
		return false
	})
}
