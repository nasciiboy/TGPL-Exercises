// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
  "fmt"
  "os"
  "net/http"
  "strings"
  "bufio"

  "golang.org/x/net/html"
)

func main() {
  for _, arg := range os.Args[1:] {
    w, i, err := CountWordsAndImages( arg )
    if err != nil {
      fmt.Fprintf( os.Stderr, "CountWordsandimages: %v\n\n", err )
      continue
    }

    fmt.Printf( "%q\nWords %d\nImgs %d\n\n", arg, w, i )
  }
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
  resp, err := http.Get(url)
  if err != nil {
    return
  }

  doc, err := html.Parse(resp.Body)
  resp.Body.Close()
  if err != nil {
    err = fmt.Errorf("parsing HTML: %s", err)
    return
  }

  words, images = countWordsAndImages(doc)
  return
}

func countWordsAndImages(n *html.Node) (words, images int) {
  if n.Type == html.ElementNode  {
    if n.Data == "script" || n.Data == "style" {
      return
    }
    if n.Data == "img" { images++ }
  }

  if n.Type == html.TextNode {
    scanner := bufio.NewScanner( strings.NewReader( n.Data ) )
    scanner.Split( bufio.ScanWords )

    for scanner.Scan() { words++ }
  }

  for c := n.FirstChild; c != nil; c = c.NextSibling {
    w, i := countWordsAndImages( c )
    words, images = words + w, images + i
  }

  return
}
