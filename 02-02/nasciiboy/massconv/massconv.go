package massconv

import "fmt"

type Pound float64
type Kilogram float64

const (
  PoundInK Kilogram = 0.453592
  KilogramInP Pound = 2.204623
)

func (p Pound)    String() string { return fmt.Sprintf("%gPd", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gKgs", k) }

//!-
