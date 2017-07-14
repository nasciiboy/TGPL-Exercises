// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
  "fmt"
  "os"
  "io"

  "golang.org/x/net/html"
)

type simpleReader struct {
  str string
  pos int
}

func (s *simpleReader) Read(p []byte) (n int, err error) {
	if s.pos >= len( s.str ) {
		return 0, io.EOF
	}

	n = copy( p, s.str[s.pos:] )
	s.pos += n
	return
}

func (s *simpleReader) Fill( w io.Reader ) (n int, err error) {
  p := make( []byte, 1024 )
  i := 0
  s.pos = 0
  for err == nil {
    i, err = w.Read( p )
    s.str += string( p[:i] )
    n += i
  }

  if err != nil && err != io.EOF { return }
  return n, nil
}

//! cat something.html | go run main.go
func main() {
  s := new(simpleReader)
  if _, err := s.Fill( os.Stdin ); err != nil {
    fmt.Fprintf( os.Stderr, "simpleReader: Fill: %v\n", err )
    os.Exit( 1 )
  }

  doc, err := html.Parse( s )
  if err != nil {
    fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
    os.Exit(1)
  }
  for _, link := range visit(nil, doc) {
    fmt.Println(link)
  }
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
  if n.Type == html.ElementNode && n.Data == "a" {
    for _, a := range n.Attr {
      if a.Key == "href" {
        links = append(links, a.Val)
      }
    }
  }

  if n.FirstChild  != nil { links = visit( links, n.FirstChild  ) }
  if n.NextSibling != nil { links = visit( links, n.NextSibling ) }

  return links
}
