package main

import (
	"log"
	"math/bits"
)

// SigSeg can represent either a set of signals, or a set of segments.
type SigSeg uint8

const (
	// Represents either signals or segments.
	SigSegA SigSeg = 1 << iota
	SigSegB SigSeg = 1 << iota
	SigSegC SigSeg = 1 << iota
	SigSegD SigSeg = 1 << iota
	SigSegE SigSeg = 1 << iota
	SigSegF SigSeg = 1 << iota
	SigSegG SigSeg = 1 << iota
)

func FromBytes(b []byte) (out SigSeg) {
	for _, c := range b {
		out.Set(c)
	}
	return
}

func (s *SigSeg) Set(v byte) {
	switch v {
	case 'a':
		*s = (*s) | SigSegA
	case 'b':
		*s = (*s) | SigSegB
	case 'c':
		*s = (*s) | SigSegC
	case 'd':
		*s = (*s) | SigSegD
	case 'e':
		*s = (*s) | SigSegE
	case 'f':
		*s = (*s) | SigSegF
	case 'g':
		*s = (*s) | SigSegG
	default:
		log.Panic("unknown signal bit", v)
	}
}

func (s SigSeg) Bits() int {
	return bits.OnesCount8(uint8(s))
}
