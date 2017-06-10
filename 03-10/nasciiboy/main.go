// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
  "fmt"
  "os"
)

func main() {
  for i := 1; i < len(os.Args); i++ {
    fmt.Printf("  %s\n", comma(os.Args[i]))
  }
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
  r := make( []byte, len(s) + len(s) / 3)

  for i, j, c := len( s ) - 1, len(r) - 1, 0; i >= 0; i, j = i - 1, j - 1 {
    r[j] = s[i]
    c++
    if c % 3 == 0 && i > 0 {
      j--
      r[j] = ','
    }
  }

  return string(r)
}
