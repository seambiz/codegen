package codegen_test

import (
	"math/big"
	"testing"

	"github.com/RoaringBitmap/roaring"
	gbit "github.com/tmthrgd/go-bitset"
	wbit "github.com/willf/bitset"
	"xojoc.pw/bitset"
)

// every role gets own bitset
// fct  4 stellig
// type 1 stellig
var startingNumber = 20000
var numberBits = 15 + startingNumber

func BenchmarkMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := map[int]struct{}{}
		for i := startingNumber; i < numberBits; i++ {
			m[i] = struct{}{}
			if _, ok := m[i]; !ok {
				b.Fail()
			}
		}
	}
}

func BenchmarkMapTest(b *testing.B) {
	m := map[int]struct{}{}
	for i := startingNumber; i < numberBits; i++ {
		m[i] = struct{}{}
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := startingNumber; i < numberBits; i++ {
			if _, ok := m[i]; !ok {
				b.Fail()
			}
		}
	}
}

func BenchmarkBig(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := big.NewInt(0)
		for i := startingNumber; i < numberBits; i++ {
			s.SetBit(s, i, 1)
			if s.Bit(i) != 1 {
				b.Fail()
			}
		}
	}
}

func BenchmarkBigTest(b *testing.B) {
	s := big.NewInt(0)
	for i := startingNumber; i < numberBits; i++ {
		s.SetBit(s, i, 1)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := startingNumber; i < numberBits; i++ {
			if s.Bit(i) != 1 {
				b.Fail()
			}
		}
	}
}

func BenchmarkBitSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := bitset.BitSet{}
		for i := startingNumber; i < numberBits; i++ {
			s.Set(i)
			if !s.Get(i) {
				b.Fail()
			}
		}
	}
}

func BenchmarkBitSetTest(b *testing.B) {
	s := bitset.BitSet{}
	for i := startingNumber; i < numberBits; i++ {
		s.Set(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := startingNumber; i < numberBits; i++ {
			if !s.Get(i) {
				b.Fail()
			}
		}
	}
}

func BenchmarkWillfBitSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := wbit.New(uint(numberBits))
		for i := startingNumber; i < numberBits; i++ {
			s.Set(uint(i))
			if !s.Test(uint(i)) {
				b.Fail()
			}
		}
	}
}

func BenchmarkWillfBitSetTest(b *testing.B) {
	s := wbit.New(uint(numberBits))
	for i := startingNumber; i < numberBits; i++ {
		s.Set(uint(i))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := startingNumber; i < numberBits; i++ {
			if !s.Test(uint(i)) {
				b.Fail()
			}
		}
	}
}

func BenchmarkRoaringBitSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := roaring.NewBitmap()
		for i := startingNumber; i < numberBits; i++ {
			s.Add(uint32(i))
			if !s.Contains(uint32(i)) {
				b.Fail()
			}
		}
	}
}

func BenchmarkRoaringBitSetTest(b *testing.B) {
	s := roaring.NewBitmap()
	for i := startingNumber; i < numberBits; i++ {
		s.Add(uint32(i))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := startingNumber; i < numberBits; i++ {
			if !s.Contains(uint32(i)) {
				b.Fail()
			}
		}
	}
}

func BenchmarkGoBitSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := gbit.New(uint(numberBits))
		for i := startingNumber; i < numberBits; i++ {
			s.Set(uint(i))
			if !s.IsSet(uint(i)) {
				b.Fail()
			}
		}
	}
}

func BenchmarkGoBitSetTest(b *testing.B) {
	s := gbit.New(uint(numberBits))
	for i := startingNumber; i < numberBits; i++ {
		s.Set(uint(i))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := startingNumber; i < numberBits; i++ {
			if !s.IsSet(uint(i)) {
				b.Fail()
			}
		}
	}
}
