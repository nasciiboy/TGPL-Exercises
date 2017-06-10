package lengthconv

func MToF(m Meter) Feet { return Feet(m / MeterInF) }

func FToM(f Feet) Meter { return Meter(f) * MeterInF }

//!-
