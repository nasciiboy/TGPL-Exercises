// Outline prints the outline of an HTML document tree.
package main

import (
  "fmt"
  "net/http"
  "os"

  "golang.org/x/net/html"
)

func main(){
  for _, url := range os.Args[1:] {
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf( os.Stderr, "http.Get: %v\n", err )
      continue
    }

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
      fmt.Fprintf( os.Stderr, "parsing %s as HTML: %v\n", url, err )
      continue
    }

    imgs := ElementsByTagName( doc, "img" )
    headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

    fmt.Printf( "============== Imgs \n" )
    PrintHtmlElements( imgs )
    fmt.Printf( "\n============== Headlines\n" )
    PrintHtmlElements( headings )
  }
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
  var nodes = make([]*html.Node, 0, 32 )

  appendNode := func(node *html.Node) {
    if node.Type == html.ElementNode {
      for _, tag := range name {
        if node.Data == tag {
          nodes = append( nodes, node )
        }
      }
    }
  }

  forEachNode(doc, appendNode, nil)

  return nodes
}

func PrintHtmlElements( nodes []*html.Node ){
  for _, node := range nodes {
    fmt.Printf( "<%s", node.Data )

    for _, attr := range node.Attr {
      fmt.Printf( " %s='%s'", attr.Key, attr.Val )
    }

    fmt.Printf( ">\n" )
  }
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
  if pre != nil {
    pre(n)
  }

  for c := n.FirstChild; c != nil; c = c.NextSibling {
    forEachNode(c, pre, post)
  }

  if post != nil {
    post(n)
  }
}
