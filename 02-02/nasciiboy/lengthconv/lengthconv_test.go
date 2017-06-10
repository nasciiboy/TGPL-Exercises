package lengthconv

import "testing"

func TestMToF( t *testing.T ){
  if f := MToF( 1 ); f != 3.280839895013123 { t.Errorf( "MToF( 1 ) != 3.280839895013123, result %g", f ) }
  if f := MToF( 3 ); f != 9.84251968503937  { t.Errorf( "MToF( 3 ) != 9.84251968503937, result %g", f ) }
}

func TestFToM( t *testing.T ){
  if m := FToM( 3 ); m != 0.9144000000000001 { t.Errorf( "FToM( 3 ) != 0.9144000000000001, result %g", m ) }
  if m := FToM( 1 ); m != 0.3048 { t.Errorf( "FToM( 1 ) != .03048 result %g", m ) }
}
