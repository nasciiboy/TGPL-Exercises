package main

import (
  "strings"
  "testing"
)

const dataEntrys = 1 << 15

var data []string
var trash string

func init() {
  for i := 0; i < dataEntrys; i++ {
    data = append( data, "abcdefg" )
  }
}

func BenchmarkEcho3( b *testing.B ){
  trash = strings.Join( data, " ")
}

func BenchmarkEcho1( b *testing.B ){
  trash, sep := "", ""
  for _, arg := range data {
    trash += sep + arg
    sep = " "
  }
}

// $ go test -bench=.
