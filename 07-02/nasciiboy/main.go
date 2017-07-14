package main

import (
  "fmt"
  "os"
  "io"
)

type byteCounter struct {
  n int64
  w io.Writer
}

func (c *byteCounter) Write(p []byte) (int, error) {
  c.n += int64( len(p) )

  var err error
  if c.w != nil {
    _, err = c.w.Write( p )
  }

  return len(p), err
}

func CountingWriter(w io.Writer) (io.Writer, *int64){
  var b byteCounter
  b.w = w

  return &b, &b.n
}

func main() {
  w, n := CountingWriter( os.Stdout )

  fmt.Fprintf( w, "hello, word!\n" )
  fmt.Printf( "writer count [%d]\n", *n )
  fmt.Fprintf( w, "1234567890\n" )
  fmt.Printf( "writer count [%d]\n", *n )
}
