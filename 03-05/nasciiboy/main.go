// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
  "image"
  "image/color"
  "image/png"
  "math/cmplx"
  "os"
  "flag"
)

var t = flag.String( "t", "m", "m Mandelbrot\n\tn Newton\n\ta Acos\n\ts Sqrt\n" )

func main() {
  flag.Parse()

  const (
    xmin, ymin, xmax, ymax = -2, -2, +2, +2
    width, height          = 1024, 1024
  )

  img := image.NewRGBA(image.Rect(0, 0, width, height))
  for py := 0; py < height; py++ {
    y := float64(py)/height*(ymax-ymin) + ymin
    for px := 0; px < width; px++ {
      x := float64(px)/width*(xmax-xmin) + xmin
      z := complex(x, y)
      // Image point (px, py) represents complex value z.
      switch *t {
      case "m": img.Set(px, py, mandelbrot(z))
      case "n": img.Set(px, py, newton(z))
      case "a": img.Set(px, py, acos(z))
      case "s": img.Set(px, py, sqrt(z))
      }
    }
  }
  png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
  const iterations = 200
  const contrast = 15

  var v complex128
  for n := 0; n < iterations; n++ {
    v = v*v + z
    if cmplx.Abs(v) > 2 {
      return MagicColor( contrast * n, iterations * contrast )
      //color.Gray{255 - contrast*n}
    }
  }
  return color.Black
}

func acos(z complex128) color.Color {
  v := cmplx.Acos(z)
  blue := uint8(real(v)*128) + 127
  red := uint8(imag(v)*128) + 127
  return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
  v := cmplx.Sqrt(z)
  blue := uint8(real(v)*128) + 127
  red := uint8(imag(v)*128) + 127
  return color.YCbCr{128, blue, red}
}

func newton(z complex128) color.Color {
  const iterations = 37
  const contrast = 7
  for i := uint8(0); i < iterations; i++ {
    z -= (z - 1/(z*z*z)) / 4
    if cmplx.Abs(z*z*z*z-1) < 1e-6 {
      return MagicColor( contrast * int(i), iterations * contrast )
    }
  }
  return color.Black
}

const colors = 255             // disponible colors
const secs   = 3               // sections of color
const sec    = colors / secs   // length of section

func MagicColor( x, scale int ) color.RGBA {
  i := x * colors / scale

  h := float64(i % sec) / sec
  s := 1.0
  l := float64(i) / colors

  r, g, b := hsl2rgb( h, s, l )

  return color.RGBA{ r, g, b, 255 }
}

func hsl2rgb( h, s, l float64 ) (red, green, blue uint8) {
  r, g, b := l, l, l

  var v float64

  if l <= 0.5 {
    v = l * (1.0 + s)
  } else {
    v = l + s - l * s
  }

  if v > 0 {
    m  := l + l - v
    sv := (v - m) / v
    h  *= 6.0
    sextant := int(h)
    fract := h - float64(sextant)
    vsf   := v * sv * fract
    mid1 := m + vsf
    mid2 := v - vsf

    switch sextant {
    case 0: r, g, b = v, mid1, m
    case 1: r, g, b = mid2, v, m
    case 2: r, g, b = m, v, mid1
    case 3: r, g, b = m, mid2, v
    case 4: r, g, b = mid1, m, v
    case 5: r, g, b = v, m, mid2
    }
  }

  red, green, blue = uint8(r * 255.0), uint8(g * 255.0), uint8(b * 255.0)
  return
}
