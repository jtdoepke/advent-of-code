package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	bitstream "github.com/dgryski/go-bitstream"
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
	p := parseInput(inp)
	var sumVersions func(*Packet) int
	sumVersions = func(p *Packet) int {
		c := int(p.Version)
		for _, sp := range p.Subpackets {
			c += sumVersions(sp)
		}
		return c
	}
	return strconv.Itoa(sumVersions(p))
}

func problemPart2(inp []byte) string {
	p := parseInput(inp)
	var evalPacket func(*Packet) int
	evalPacket = func(p *Packet) int {
		switch p.Type {
		case PacketTypeSum:
			sum := 0
			for _, sp := range p.Subpackets {
				sum += evalPacket(sp)
			}
			return sum

		case PacketTypeProduct:
			product := 1
			for _, sp := range p.Subpackets {
				product *= evalPacket(sp)
			}
			return product

		case PacketTypeMinimum:
			min := math.MaxInt
			for _, sp := range p.Subpackets {
				if v := evalPacket(sp); v < min {
					min = v
				}
			}
			return min

		case PacketTypeMaximum:
			max := math.MinInt
			for _, sp := range p.Subpackets {
				if v := evalPacket(sp); v > max {
					max = v
				}
			}
			return max

		case PacketTypeLiteral:
			return int(p.Literal)

		case PacketTypeGreaterThan:
			if evalPacket(p.Subpackets[0]) > evalPacket(p.Subpackets[1]) {
				return 1
			} else {
				return 0
			}

		case PacketTypeLessThan:
			if evalPacket(p.Subpackets[0]) < evalPacket(p.Subpackets[1]) {
				return 1
			} else {
				return 0
			}

		case PacketTypeEqualTo:
			if evalPacket(p.Subpackets[0]) == evalPacket(p.Subpackets[1]) {
				return 1
			} else {
				return 0
			}
		}
		log.Panic("unknown packet type", p.Type)
		return 0
	}

	return strconv.Itoa(evalPacket(p))
}

func parseInput(inp []byte) *Packet {
	r, err := hexToBitstream(bytes.TrimSpace(inp))
	if err != nil {
		log.Panic(err)
	}
	_, p, err := readPacket(r)
	if err != nil {
		log.Panic(err)
	}
	return p
}

type PacketType byte

const (
	PacketTypeSum PacketType = iota
	PacketTypeProduct
	PacketTypeMinimum
	PacketTypeMaximum
	PacketTypeLiteral
	PacketTypeGreaterThan
	PacketTypeLessThan
	PacketTypeEqualTo
)

type PacketLengthType byte

const (
	PacketLengthTypeTotal      PacketLengthType = 0
	PacketLengthTypeSubpackets PacketLengthType = 1
)

type Packet struct {
	Version byte
	Type    PacketType

	Literal    uint64
	Subpackets []*Packet
}

func readPacket(r *bitstream.BitReader) (int, *Packet, error) {
	read := 0

	version, err := r.ReadBits(3)
	if err != nil {
		return 0, nil, err
	}
	read += 3

	pTypeInt, err := r.ReadBits(3)
	if err != nil {
		return 0, nil, err
	}
	pType := PacketType(pTypeInt)
	read += 3

	p := &Packet{
		Version: byte(version),
		Type:    pType,
	}

	switch pType {
	case PacketTypeLiteral:
		n, v, err := readPacketLiteral(r)
		if err != nil {
			return 0, nil, err
		}
		p.Literal = v
		read += n
	default: // Operator packet
		n, subpackets, err := readSubpackets(r)
		if err != nil {
			return 0, nil, err
		}
		p.Subpackets = subpackets
		read += n
	}

	return read, p, nil
}

func readPacketLiteral(r *bitstream.BitReader) (int, uint64, error) {
	read := 0

	doRead := func() (stop bool, v uint64, err error) {
		bits, err := r.ReadBits(5)
		if err != nil {
			return true, 0, err
		}
		cont := bits & 0b10000
		value := bits & 0b1111
		read += 5
		return (cont == 0), value, nil
	}

	stop, v, err := doRead()
	if err != nil {
		return 0, 0, err
	}
	finalValue := v
	for !stop {
		finalValue = finalValue << 4
		stop, v, err = doRead()
		if err != nil {
			return 0, 0, err
		}
		finalValue = finalValue | v
	}
	return read, finalValue, nil
}

func readSubpackets(r *bitstream.BitReader) (int, []*Packet, error) {
	var subpackets []*Packet
	read := 0
	v, err := r.ReadBits(1)
	if err != nil {
		return 0, nil, err
	}
	read++
	lt := PacketLengthType(v)
	switch lt {
	case PacketLengthTypeTotal:
		v, err := r.ReadBits(15)
		if err != nil {
			return 0, nil, err
		}
		pLen := int(v)
		read += 15
		subread := 0
		for subread < pLen {
			n, sp, err := readPacket(r)
			if err != nil {
				return 0, nil, err
			}
			subpackets = append(subpackets, sp)
			subread += n
		}
		read += subread

	case PacketLengthTypeSubpackets:
		v, err := r.ReadBits(11)
		if err != nil {
			return 0, nil, err
		}
		pLen := int(v)
		read += 11
		for i := 0; i < pLen; i++ {
			n, sp, err := readPacket(r)
			if err != nil {
				return 0, nil, err
			}
			subpackets = append(subpackets, sp)
			read += n
		}
	}

	return read, subpackets, nil
}

func hexToBitstream(data []byte) (*bitstream.BitReader, error) {
	b, err := hex.DecodeString(string(bytes.TrimSpace(data)))
	if err != nil {
		return nil, err
	}
	return bitstream.NewReader(bytes.NewReader(b)), nil
}
