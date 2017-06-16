// Issues prints a table of GitHub issues matching the search terms.
package main

import (
  "flag"
  "fmt"
  "path"

	"./xkcd"
)

var t = flag.Bool( "t", false, "include transcription" )

func main(){
  flag.Parse()

  if !xkcd.ExistData() {
    xkcd.GetAllXKCD()
    xkcd.MakeDB()
  }

  xkcd.LoadDB()

  for _, arg := range flag.Args() {
    for _, comic := range xkcd.Search( arg ) {
      ext := path.Ext( comic.ImageURL )
      fmt.Printf( "[%04d] %q\n\t%s ./img/%04d%s\n", comic.Num, comic.Title, comic.ImageURL, comic.Num, ext )
      if *t {
        fmt.Println( "--------------------------------------------------------------------------------" )
        fmt.Printf( "%s\n", comic.Transcript )
      }
      fmt.Println( "--------------------------------------------------------------------------------" )
    }
  }
}
