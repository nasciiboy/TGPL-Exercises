// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
  "fmt"
  "crypto/sha256"
)

func main() {
  c1 := sha256.Sum256([]byte("x"))
  c2 := sha256.Sum256([]byte("X"))

  fmt.Printf("%x\n%x\n%t\ndiffBits() %d\n", c1, c2, c1 == c2, diffBits(&c1, &c2))
}

func diffBits( a, b *[sha256.Size]byte ) (n int) {
  for i := range a {
    for j := uint(1); j < 256; j <<= 1 {
      if a[i] & byte(j) != b[i] & byte(j) { n++ }
    }
  }

  return
}

//!-
// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
// false
// diffBits() 125
