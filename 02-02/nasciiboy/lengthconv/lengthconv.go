package lengthconv

import "fmt"

type Feet float64
type Meter float64

const (
  MeterInF Meter = 0.3048
  FeetInM  Feet  = 3.280839895013123
)

func (f Feet)  String() string { return fmt.Sprintf("%gf", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

//!-
