package main

import (
  "fmt"
  "os"
  "strconv"
)

func main(){
  for _, arg := range os.Args[1:] {
    n, err := strconv.Atoi( arg )
    if err != nil {
      fmt.Fprintf( os.Stderr, "Garbage entry %s: %v\n", arg, err )
      continue
    }

    fmt.Printf( "square( %3d ) == %6d\n", n, square( n ) )
  }
}

func square( n int ) (r int){
  defer func(){
    recover()
    r = n * n
  }()
  panic( nil )
}

//!- output
//  go run main.go 5 10 15 100 "hola amigo" 9 3
//  square(   5 ) ==     25
//  square(  10 ) ==    100
//  square(  15 ) ==    225
//  square( 100 ) ==  10000
//  Garbage entry hola: strconv.Atoi: parsing "hola amigo": invalid syntax
//  square(   9 ) ==     81
//  square(   3 ) ==      9
