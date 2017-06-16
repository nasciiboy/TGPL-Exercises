package main

import (
  "bufio"
  "sort"
  "fmt"
  "os"
)

type count struct {
  word string
  n    int
}

type byNum []count

func (c byNum) Len() int           { return len(c) }
func (c byNum) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c byNum) Less(i, j int) bool {
  if c[i].n < c[j].n {
    return true
  }
  return false
}

func main() {
  words   := make(map[string]int)
  scanner := bufio.NewScanner( os.Stdin )
  scanner.Split( bufio.ScanWords )

  for scanner.Scan() {
    words[ scanner.Text() ]++
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintf( os.Stderr, "wordfreq: %v\n", err )
    os.Exit( 2 )
  }

  sw, i := make( byNum, len(words) ), 0
  for word, n := range words {
    sw[i].word = word
    sw[i].n    = n
    i++
  }

  sort.Sort(sw)

  for i := range sw {
    fmt.Printf( "[%5d] %q\n", sw[i].n, sw[i].word )
  }
}

//!-
