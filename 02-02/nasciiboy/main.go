package main

import (
  "fmt"
  "os"
  "strconv"
  "flag"
  "bufio"

  "./tempconv"
  "./massconv"
  "./lengthconv"
)

var t = flag.Bool( "t", false, "conv format to temp" )
var m = flag.Bool( "m", false, "conv format to mass" )
var l = flag.Bool( "l", false, "conv format to lenght" )

func main() {
  flag.Parse()

  if !*t && !*m && !*l {
    fmt.Fprintf( os.Stderr, "conv: set format, --help to see formats\n" )
    os.Exit(1)
  }

  if len( flag.Args() ) == 0 {
    in := bufio.NewScanner( os.Stdin )
    for in.Scan() {
      v, err := strconv.ParseFloat(in.Text(), 64)
      if err != nil {
        fmt.Fprintf( os.Stderr, "conv: %v\n", err )
        os.Exit(1)
      }

      if *t { toTemp( v ) }
      if *m { toMass( v ) }
      if *l { toLeng( v ) }
    }

    return
  }

  for _, arg := range flag.Args() {
    v, err := strconv.ParseFloat(arg, 64)
    if err != nil {
      fmt.Fprintf( os.Stderr, "conv: %v\n", err )
      os.Exit(1)
    }

    if *t { toTemp( v ) }
    if *m { toMass( v ) }
    if *l { toLeng( v ) }
  }
}

func toTemp( t float64 ){
  f := tempconv.Fahrenheit(t)
  c := tempconv.Celsius(t)
  k := tempconv.Kelvin(t)

  fmt.Printf("%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
    c, tempconv.CToF(c),
    c, tempconv.CToK(c),
    f, tempconv.FToC(f),
    f, tempconv.FToK(f),
    k, tempconv.KToC(k),
    k, tempconv.KToF(k))
}

func toMass( t float64 ){
  p := massconv.Pound(t)
  k := massconv.Kilogram(t)

  fmt.Printf("%s = %s\n%s = %s\n",
    p, massconv.PToK(p),
    k, massconv.KToP(k))
}

func toLeng( t float64 ){
  f := lengthconv.Feet(t)
  m := lengthconv.Meter(t)

  fmt.Printf("%s = %s\n%s = %s\n",
    f, lengthconv.FToM(f),
    m, lengthconv.MToF(m))
}
