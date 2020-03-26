// Comma prints its argument numbers with a comma at each power of 1000.
// Example:
// 	$ go build main.go
//	$ ./main -1 +12 -123 +1234 -12345 +123456 -1234567 +1234567890 -12345678901234
//  -1
//  +12
//  -123
//  +1,234
//  -12,345
//  +123,456
//  -1,234,567
//  +1,234,567,890
//  -12,345,678,901,234
// 	$ go build main.go
//	$ ./main .0 .01 .012 1.0123 -1.99456 +77.070707 -1234567.7707
//  .0
//  .01
//  .012
//  1.0123
//  -1.99456
//  +77.070707
//  -1,234,567.7707
package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  for i := 1; i < len(os.Args); i++ {
    fmt.Printf("  %s\n", commaDecimal(os.Args[i]))
  }
}

func commaDecimal( s string ) string {
  sign, decimal := "", ""
  if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
    sign = s[:1]
    s = s[1:]
  }

  if dot := strings.LastIndex(s, "."); dot >= 0 {
    decimal = s[dot:]
    s       = s[:dot]
  }

  return sign + comma( s ) + decimal
}

// comma inserts commas in a non-negative decimal integer string.
func comma( s string ) string {
  r := make( []byte, len(s) + len(s) / 3)

  for i, j, c := len(s) - 1, len(r) - 1, 0; i >= 0; i, j = i - 1, j - 1 {
    r[j] = s[i]
    c++
    if c % 3 == 0 && i > 0 {
      j--
      r[j] = ','
    }
  }

  if len(r) > 0 && r[0] == 0 { return string(r[1:]) }
  return string(r)
}
