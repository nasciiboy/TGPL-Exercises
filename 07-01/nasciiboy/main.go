package main

import (
  "bytes"
  "bufio"
  "fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
  *c += ByteCounter(len(p)) // convert int to ByteCounter
  return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
  scanner := bufio.NewScanner( bytes.NewReader( p ) )
  scanner.Split( bufio.ScanWords )

  n := 0
  for scanner.Scan() { n++ }
  *w += WordCounter(n)

  return n, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
  scanner := bufio.NewScanner( bytes.NewReader( p ) )
  scanner.Split( bufio.ScanLines )

  n := 0
  for scanner.Scan() { n++ }
  *l += LineCounter( n )

  return n, nil
}

func main() {
  var c ByteCounter
  c.Write([]byte("hello"))
  fmt.Println(c) // "5", = len("hello")

  c = 0 // reset the counter
  var name = "Dolly"
  fmt.Fprintf(&c, "hello, %s", name)
  fmt.Println(c) // "12", = len("hello, Dolly")

  var w WordCounter
  w.Write([]byte("hello hillo hollo hullo")) // 4
  fmt.Println(w) // "4"

  w = 0
  name = "1 2 3 4 1969 778"
  fmt.Fprintf(&w, "hello, %s", name)
  fmt.Println(w) // 7

  var l LineCounter
  l.Write([]byte("hello\n hillo\n hollo\n hullo")) // 4
  fmt.Println(l) // "4"

  l = 0
  name = "1\n 2\n 3\n 4\n 1969\n 778"
  fmt.Fprintf(&l, "hello\n %s", name)
  fmt.Println(l) // 7
}
