package main

import "fmt"

func eliminateAdjacentDuplicates( data []string ) []string {
  if len( data ) == 0 { return data }

  prev := 0
  for next := 1; next < len(data); next++ {
    if data[prev] != data[next] {
      prev++
      data[prev] = data[next]
    }
  }

  return data[:prev + 1]
}

func main(){
  data := []string{
    "start",
    "start",
    "start",
    "line 1",
    "line 2",
    "line 2",
    "line 2",
    "line 2",
    "line 3",
    "end",
    "end",
  }

  data = eliminateAdjacentDuplicates( data )

  for i := range data {
    fmt.Printf( "%s\n", data[i] )
  }
}
