// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
  "image"
  "image/color"
  "image/gif"
  "io"
  "math"
  "math/rand"
  "strconv"
)

import (
  "log"
  "net/http"
  "time"
)

var palette = []color.Color{color.White, color.Black}

const (
  whiteIndex = 0 // first color in palette
  blackIndex = 1 // next color in palette
)

type lconfig struct {
  cycles float64 // number of complete x oscillator revolutions
  res    float64 // angular resolution
  freq   float64 // relative frequency of y oscillator
  size   int     // image canvas covers [-size..+size]
  frames int     // number of animation frames
  delay  int     // delay between frames in 10ms units
}

func main() {
  rand.Seed(time.Now().UnixNano())
  lconf := lconfig {
    cycles  : 5,
    res     : 0.001,
    freq    : rand.Float64() * 3.0,
    size    : 100,
    frames  : 64,
    delay   : 8,
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    confs := r.URL.Query()
    for i, c := range confs {
      switch i {
      case "cycles" : lconf.cycles, _ = strconv.ParseFloat( c[0], 64 )
      case "freq"   : lconf.freq  , _ = strconv.ParseFloat( c[0], 64 )
      case "res"    : lconf.res   , _ = strconv.ParseFloat( c[0], 64 )
      case "size"   : lconf.size  , _ = strconv.Atoi( c[0] )
      case "frames" : lconf.frames, _ = strconv.Atoi( c[0] )
      case "delay"  : lconf.delay , _ = strconv.Atoi( c[0] )
      }
    }

    lissajous(w, lconf)
  })

  log.Fatal(http.ListenAndServe("localhost:8000", nil))
  return
}

func lissajous(out io.Writer, set lconfig) {
  anim  := gif.GIF{LoopCount: set.frames}
  phase := 0.0 // phase difference
  for i := 0; i < set.frames; i++ {
    rect := image.Rect(0, 0, 2*set.size+1, 2*set.size+1)
    img := image.NewPaletted(rect, palette)
    for t := 0.0; t < set.cycles*2*math.Pi; t += set.res {
      x := math.Sin(t)
      y := math.Sin(t*set.freq + phase)
      img.SetColorIndex(set.size+int(x*float64(set.size)+0.5), set.size+int(y*float64(set.size)+0.5), blackIndex)
    }
    phase += 0.1
    anim.Delay = append(anim.Delay, set.delay)
    anim.Image = append(anim.Image, img)
  }
  gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
