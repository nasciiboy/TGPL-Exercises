// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
  "strings"
)

const prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
    if strings.HasPrefix( url, prefix ) == false {
      url = prefix + url
    }

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy( os.Stdout, resp.Body)
    resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

    fmt.Printf( "Status %s\n", resp.Status )
	}
}
