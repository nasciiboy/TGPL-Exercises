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
      s[j] = ' '
      j++
    }

    w := utf8.RuneLen( r )
    copy( s[j:], s[i:i+w] )
    j += w
  }

  return s[:j]
}

func main(){
  a := []byte("    a \t\vspaç£ t£est \n to")
  fmt.Printf( "%q\n", string(a) ) // "    a \t\vspaç£ t£est \n to"
  a = rmSpaces( a )
  fmt.Printf( "%q\n", string(a) ) // " a spaç£ t£est to"
}
