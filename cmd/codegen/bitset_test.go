package main

import (
	"math/big"
	"testing"

	wbit "github.com/willf/bitset"
	"xojoc.pw/bitset"
)

var numberBits = 10

func BenchmarkBig(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := big.NewInt(0)
		for i := 0; i < numberBits; i++ {
			s.SetBit(s, i, 1)
			if s.Bit(i) != 1 {
				b.Fail()
			}
		}
	}
}

func BenchmarkBitSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := bitset.BitSet{}
		for i := 0; i < numberBits; i++ {
			s.Set(i)
			if !s.Get(i) {
				b.Fail()
			}
		}
	}
}

func BenchmarkWillfBitSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := wbit.New(uint(numberBits))
		for i := 0; i < numberBits; i++ {
			s.Set(uint(i))
			if !s.Test(uint(i)) {
				b.Fail()
			}
		}
	}
}
