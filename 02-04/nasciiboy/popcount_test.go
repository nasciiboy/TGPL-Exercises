// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount

import "testing"

// -- Alternative implementations --

func BitCount(x uint64) int {
  // Hacker's Delight, Figure 5-2.
  x = x - ((x >> 1) & 0x5555555555555555)
  x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
  x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
  x = x + (x >> 8)
  x = x + (x >> 16)
  x = x + (x >> 32)
  return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
  n := 0
  for x != 0 {
    x = x & (x - 1) // clear rightmost non-zero bit
    n++
  }
  return n
}

func PopCountByShifting(x uint64) int {
  n := 0
  for i := uint(0); i < 64; i++ {
    if x&(1<<i) != 0 {
      n++
    }
  }
  return n
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
  for i := 0; i < b.N; i++ {
    PopCount(0x1234567890ABCDEF)
  }
}

func BenchmarkPopCountByLoop(b *testing.B) {
  for i := 0; i < b.N; i++ {
    PopCountByLoop(0x1234567890ABCDEF)
  }
}

func BenchmarkPopCountByRightMostBit(b *testing.B) {
  for i := 0; i < b.N; i++ {
    PopCountByRightMostBit(0x1234567890ABCDEF)
  }
}

func BenchmarkBitCount(b *testing.B) {
  for i := 0; i < b.N; i++ {
    BitCount(0x1234567890ABCDEF)
  }
}

func BenchmarkPopCountByClearing(b *testing.B) {
  for i := 0; i < b.N; i++ {
    PopCountByClearing(0x1234567890ABCDEF)
  }
}

func BenchmarkPopCountByShifting(b *testing.B) {
  for i := 0; i < b.N; i++ {
    PopCountByShifting(0x1234567890ABCDEF)
  }
}

// Go 1.8, 1.2GHz Celeron 847
// $ go test -bench=.
// BenchmarkPopCount-2                     2000000000               0.98 ns/op
// BenchmarkPopCountByLoop-2               20000000                74.2 ns/op
// BenchmarkPopCountByRightMostBit-2        5000000               256 ns/op
// BenchmarkBitCount-2                     2000000000               0.95 ns/op
// BenchmarkPopCountByClearing-2           20000000                96.0 ns/op
// BenchmarkPopCountByShifting-2            5000000               255 ns/op
