package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	bitstream "github.com/dgryski/go-bitstream"
	"github.com/stretchr/testify/assert"
)

func getTestData(t *testing.T) []byte {
	f, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func Test_problemPart1(t *testing.T) {
	tests := []struct {
		name string
		inp  string
		want string
	}{
		{
			name: "16",
			inp:  "8A004A801A8002F478",
			want: "16",
		},
		{
			name: "12",
			inp:  "620080001611562C8802118E34",
			want: "12",
		},
		{
			name: "23",
			inp:  "C0015000016115A2E0802F182340",
			want: "23",
		},
		{
			name: "31",
			inp:  "A0016C880162017C3686B18A3D4780",
			want: "31",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := problemPart1([]byte(tt.inp))
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_problemPart2(t *testing.T) {
	tests := []struct {
		name string
		inp  string
		want string
	}{
		{
			name: "sum",
			inp:  "C200B40A82",
			want: "3",
		},
		{
			name: "product",
			inp:  "04005AC33890",
			want: "54",
		},
		{
			name: "min",
			inp:  "880086C3E88112",
			want: "7",
		},
		{
			name: "max",
			inp:  "CE00C43D881120",
			want: "9",
		},
		{
			name: "lt",
			inp:  "D8005AC2A8F0",
			want: "1",
		},
		{
			name: "gt",
			inp:  "F600BC2D8F",
			want: "0",
		},
		{
			name: "eq",
			inp:  "9C005AC2F8F0",
			want: "0",
		},
		{
			name: "complex",
			inp:  "9C0141080250320F1802104A08",
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := problemPart2([]byte(tt.inp))
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_hexToBitstream(t *testing.T) {
	bs, err := hexToBitstream([]byte("D2FE28"))
	if !assert.NoError(t, err) {
		return
	}
	got := strings.Builder{}
	for {
		b, err := bs.ReadBit()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		if b {
			got.WriteByte('1')
		} else {
			got.WriteByte('0')
		}
	}
	assert.Equal(t, "110100101111111000101000", got.String())
}

func Test_readPacketLiteral(t *testing.T) {
	r := newBitstreamReader(t, 0b101111111000101000, 18)

	n, got, err := readPacketLiteral(r)
	if assert.NoError(t, err) {
		assert.Equal(t, 15, n)
		assert.EqualValues(t, 2021, got)
	}
}

func newBitstreamReader(t *testing.T, u uint64, bits int) *bitstream.BitReader {
	buf := bytes.Buffer{}
	w := bitstream.NewWriter(&buf)
	err := w.WriteBits(u, bits)
	if err != nil {
		t.Fatal(err)
	}
	r := bitstream.NewReader(bytes.NewReader(buf.Bytes()))
	return r
}

func Test_readSubpackets(t *testing.T) {
	type args struct {
		r *bitstream.BitReader
	}
	tests := []struct {
		name           string
		args           args
		wantN          int
		wantSubpackets []*Packet
		wantErr        bool
	}{
		{
			name: "example",
			args: args{
				r: newBitstreamReader(t, 0b10000000001101010000001100100000100011000001100000, 50),
			},
			wantN: 45,
			wantSubpackets: []*Packet{
				{
					Version: 2,
					Type:    PacketTypeLiteral,
					Literal: 1,
				},
				{
					Version: 4,
					Type:    PacketTypeLiteral,
					Literal: 2,
				},
				{
					Version: 1,
					Type:    PacketTypeLiteral,
					Literal: 3,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, gotSubpackets, err := readSubpackets(tt.args.r)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantN, gotN)
			assert.Equal(t, tt.wantSubpackets, gotSubpackets)
		})
	}
}
