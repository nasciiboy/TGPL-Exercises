package main

import (
  "fmt"
)

func reverse(s []byte) {
  r := []rune(string(s))
  for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
  }

  copy( s, []byte(string(r)) )
}

func main(){
  data := [][]byte{
    []byte( "start" ),
    []byte( "tert" ),
    []byte( "12Ø34" ),
    []byte( "123Ø4" ),
    []byte( "1Ø234" ),
    []byte( "¤2Ø34" ),
    []byte( "1¤3Ø4" ),
    []byte( "1Ø2¤4" ),
    []byte( "Øe¤¥næn" ),
  }

  for i := range data {
    fmt.Printf( "%s | ", data[i] )
    reverse( data[i] )
    fmt.Printf( "%s | ", data[i] )
    reverse( data[i] )
    fmt.Printf( "%s\n", data[i] )
  }
}
