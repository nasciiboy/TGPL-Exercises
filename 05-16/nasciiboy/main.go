package main

import "fmt"

func main(){
  data := []string{
    "Sonic Youth",
    "Confusion Is Sex",
    "Kill Yr Idols",
    "Bad Moon Rising",
    "Sister",
    "Hold That Tiger",
    "Daydream Nation",
    "Goo",
    "100%",
    "Dirty",
    "Experimental Jet Set, Trash and No Star",
    "Washing Machine",
    "A Thousand Leaves",
    "NYC Ghosts & Flowers",
    "Murry Street",
    "Sonic Nurse",
    "Rather Ripped",
    "The Eternal",
  }

  fmt.Println( Join( "\n", data... ) )
}

// from github.com/golang/go/src/strings/string.go
func Join( sep string, a ...string ) string {
  switch len(a) {
  case 0: return ""
  case 1: return a[0]
  case 2: return a[0] + sep + a[1]
  case 3: return a[0] + sep + a[1] + sep + a[2]
  }

  n := len(sep) * (len(a) - 1)
  for i := 0; i < len(a); i++ {
    n += len(a[i])
  }

  b := make([]byte, n)
  bp := copy(b, a[0])
  for _, s := range a[1:] {
    bp += copy(b[bp:], sep)
    bp += copy(b[bp:], s)
  }
  return string(b)
}
