package main

import "fmt"

func rotate( s []int, n int ){
  w := len( s )
  if n < w && n > 0 {
    start := make( []int, n )
    copy( start, s[:n] )
    copy( s, s[n:] )
    copy( s[w-n:], start )
  }
}

func main(){
  a := [...]int{0, 1, 2, 3, 4, 5, 6}
  rotate( a[:], 2 )
  fmt.Println( a ) // "[2 3 4 5 6 0 1]"
}
