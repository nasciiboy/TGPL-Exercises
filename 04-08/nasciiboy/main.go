// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "unicode"
  "unicode/utf8"
)

func main() {
  counts := make(map[rune]int)    // counts of Unicode characters
  var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
  invalid := 0                    // count of invalid UTF-8 characters

  var isControl, isDigit, isGraphic, isLetter, isLower, isMark, isNumber, isPrint, isPunct, isSpace, isSymbol, isTitle, isUpper int

  in := bufio.NewReader(os.Stdin)
  for {
    r, n, err := in.ReadRune() // returns rune, nbytes, error
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
      os.Exit(1)
    }
    if r == unicode.ReplacementChar && n == 1 {
      invalid++
      continue
    }
    counts[r]++
    utflen[n]++

    if unicode.IsControl( r ) { isControl++ }
    if unicode.IsDigit  ( r ) { isDigit++   }
    if unicode.IsGraphic( r ) { isGraphic++ }
    if unicode.IsLetter ( r ) { isLetter++  }
    if unicode.IsLower  ( r ) { isLower++   }
    if unicode.IsMark   ( r ) { isMark++    }
    if unicode.IsNumber ( r ) { isNumber++  }
    if unicode.IsPrint  ( r ) { isPrint++   }
    if unicode.IsPunct  ( r ) { isPunct++   }
    if unicode.IsSpace  ( r ) { isSpace++   }
    if unicode.IsSymbol ( r ) { isSymbol++  }
    if unicode.IsTitle  ( r ) { isTitle++   }
    if unicode.IsUpper  ( r ) { isUpper++   }

  }
  fmt.Printf("rune\tcount\n")
  for c, n := range counts {
    fmt.Printf("%q\t%d\n", c, n)
  }
  fmt.Print("\nlen\tcount\n")
  for i, n := range utflen {
    if i > 0 {
      fmt.Printf("%d\t%d\n", i, n)
    }
  }
  if invalid > 0 {
    fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
  }

  if isGraphic > 0 { fmt.Printf( "\n%8d\tGraphic\tcharacters\n", isGraphic ) }
  if isPrint   > 0 { fmt.Printf( "%8d\tPrint\tcharacters\n",   isPrint   ) }
  if isLetter  > 0 { fmt.Printf( "%8d\tLetter\tcharacters\n",  isLetter  ) }
  if isLower   > 0 { fmt.Printf( "%8d\tLower\tcharacters\n",   isLower   ) }
  if isSpace   > 0 { fmt.Printf( "%8d\tSpace\tcharacters\n",   isSpace   ) }
  if isSymbol  > 0 { fmt.Printf( "%8d\tSymbol\tcharacters\n",  isSymbol  ) }
  if isPunct   > 0 { fmt.Printf( "%8d\tPunct\tcharacters\n",   isPunct   ) }
  if isUpper   > 0 { fmt.Printf( "%8d\tUpper\tcharacters\n",   isUpper   ) }
  if isControl > 0 { fmt.Printf( "%8d\tControl\tcharacters\n", isControl ) }
  if isMark    > 0 { fmt.Printf( "%8d\tMark\tcharacters\n",    isMark    ) }
  if isNumber  > 0 { fmt.Printf( "%8d\tNumber\tcharacters\n",  isNumber  ) }
  if isDigit   > 0 { fmt.Printf( "%8d\tDigit\tcharacters\n",   isDigit   ) }
  if isTitle   > 0 { fmt.Printf( "%8d\tTitle\tcharacters\n",   isTitle   ) }
}

//!-
