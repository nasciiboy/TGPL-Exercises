package main

import (
  "bufio"
  "fmt"
  "os"
)

func main(){
  counts := make(map[string]map[string]int) // [line][filename]n
  files := os.Args[1:]

  if len(files) == 0 {
    countLines( os.Stdin, counts )
  } else {
    for _, arg := range files {
      f, err := os.Open( arg )
      if err != nil {
        fmt.Fprintf( os.Stderr, "dup2: %v\n", err )
        continue
      }
      countLines( f, counts )
      f.Close()
    }
  }

  for line, fileMap := range counts {
    tot := 0
    for _, n := range fileMap { tot += n }

    if tot > 1 {
      fmt.Printf("%d %s\n", tot, line)

      for file, n := range fileMap {
        fmt.Printf( "  %d %s\n", n, file )
      }
    }
  }
}

func countLines(f *os.File, counts map[string]map[string]int) {
  input := bufio.NewScanner(f)
  for input.Scan() {
    line := input.Text()

    if counts[ line ] == nil {
      m := make(map[string]int)
      counts[ line ] = m
    }

    counts[line][f.Name()]++
  }

  // NOTE: ignoring potential errors from input.Err()
}

//!- $ go run main.go a.txt b.txt
