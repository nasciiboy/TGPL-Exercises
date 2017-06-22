// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
  "fmt"
  "os"

  "golang.org/x/net/html"
)

func main() {
  doc, err := html.Parse(os.Stdin)
  if err != nil {
    fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
    os.Exit(1)
  }

  pPrint( doc )
}

// visit appends to links each link found in n and returns the result.
func pPrint( n *html.Node ) {
  if n.Type == html.ElementNode  {
    if n.Data == "script" || n.Data == "style" {
      return
    }
  }

  if n.Type == html.TextNode {
    fmt.Printf( "%s", n.Data )
  }

  for c := n.FirstChild; c != nil; c = c.NextSibling {
    pPrint(c)
  }
}
