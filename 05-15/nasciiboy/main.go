package main

import "fmt"

func main(){
  data := []int{ -5, -50, 0, 5, 60, 50 }

  fmt.Printf( "%v max() %d\n", data, max( data... ) )
  fmt.Printf( "%v max() %d\n", data[0], max( data[0] ) )

  fmt.Printf( "%v Max() %d\n", data, Max( data[0], data[1:]... ) )
  fmt.Printf( "%v Max() %d\n", data[0], Max( data[0] ) )

  fmt.Printf( "%v min() %d\n", data, min( data... ) )
  fmt.Printf( "%v min() %d\n", data[0], min( data[0] ) )

  fmt.Printf( "%v Min() %d\n", data, Min( data... ) )
  fmt.Printf( "%v Min() %d\n", data[0], Min( data[0] ) )
}

func max( data ...int ) (m int) {
  m = data[0]
  for _, c := range data[1:] {
    if m < c { m = c }
  }
  return
}

func Max( m int, data ...int ) int {
  for _, c := range data {
    if m < c { m = c }
  }
  return m
}

func min( data ...int ) (m int) {
  m = data[0]
  for _, c := range data[1:] {
    if m > c { m = c }
  }
  return
}

func Min( data ...int ) (m int) {
  if len( data ) < 1 {
    panic( "Min() need minimum one argument, ANTA BAKA!" )
  }

  m = data[0]
  for _, c := range data {
    if m > c { m = c }
  }
  return m
}
