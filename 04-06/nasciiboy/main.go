package main

import (
  "fmt"
  "unicode"
  "unicode/utf8"
)

func rmSpaces( s []byte ) []byte {
  inspace, j := false, 0
  for i, r := range string(s) {
    if unicode.IsSpace( r ) {
      inspace = true
      continue
    } else if inspace {
      inspace = false
      r = ' '
    }

    w := utf8.RuneLen( r )
    copy( s[j:], s[i:i+w] )
    j += w
  }

  return s[:j]
}

func main(){
  a := []byte("u\t n\v  \t\ni c      o᠎᠎᠎᠎᠎᠎‍d　e")
  a = rmSpaces( a )
  fmt.Println( string(a) ) // "unicode"
}
