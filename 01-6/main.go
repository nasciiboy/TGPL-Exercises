// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
  "image"
  "image/color"
  "image/gif"
  "io"
  "math"
  "math/rand"
  "os"
  "time"
)

const nColors = 10 // >= 3

func main() {
  seed := time.Now()
  rand.Seed( seed.Unix() )

  var palette []color.Color

  for i := 0; i < nColors; i++ {
    r := uint8(rand.Uint32() % 256)
    g := uint8(rand.Uint32() % 256)
    b := uint8(rand.Uint32() % 256)
    palette = append( palette, color.RGBA{ r, g, b, 0xFF} )
  }

  lissajous(os.Stdout, palette)
}

func lissajous(out io.Writer, palette []color.Color) {
  const (
    cycles  = 5     // number of complete x oscillator revolutions
    res     = 0.001 // angular resolution
    size    = 100   // image canvas covers [-size..+size]
    nframes = 64    // number of animation frames
    delay   = 8     // delay between frames in 10ms units
  )
  freq := rand.Float64() * 3.0 // relative frequency of y oscillator
  anim := gif.GIF{LoopCount: nframes}
  phase := 0.0                 // phase difference
  for i := 0; i < nframes; i++ {
    rect   := image.Rect(0, 0, 2*size+1, 2*size+1)
    img    := image.NewPaletted(rect, palette)
    ncolor := uint8(i % len(palette))
    for t := 0.0; t < cycles*2*math.Pi; t += res {
      x := math.Sin(t)
      y := math.Sin(t*freq + phase)
      img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), ncolor )
    }
    phase += 0.1
    anim.Delay = append(anim.Delay, delay)
    anim.Image = append(anim.Image, img)
  }
  gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
