package tempconv

import "testing"

func TestCToX( t *testing.T ){
  if f := CToF(   0 ); f != 32     { t.Errorf( "CToF(   0 ) != 32, result %g", f ) }
  if f := CToF( 100 ); f != 212    { t.Errorf( "CToF( 100 ) != 32, result %g", f ) }
  if k := CToK(   0 ); k != 273.15 { t.Errorf( "CToK(   0 ) != 273.15, result %g", k ) }
  if k := CToK( 100 ); k != 373.15 { t.Errorf( "CToK( 100 ) != 373.15, result %g", k ) }
}

func TestFToX( t *testing.T ){
  if c := FToC(  32 ); c !=   0    { t.Errorf( "FToC(  32 ) != 0, result %g", c ) }
  if c := FToC( 212 ); c != 100    { t.Errorf( "FToC( 212 ) != 100, result %g", c ) }
  if k := FToK(  32 ); k != 273.15 { t.Errorf( "FToK(  32 ) != 273.15, result %g", k ) }
  if k := FToK( 212 ); k != 373.15 { t.Errorf( "FToK( 212 ) != 373.15, result %g", k ) }
}

func TestKToX( t *testing.T ){
  if c := KToC( 273.15 ); c !=   0 { t.Errorf( "KToC( 273.15 ) != 0, result %g", c ) }
  if c := KToC( 373.15 ); c != 100 { t.Errorf( "KToC( 373.15 ) != 100, result %g", c ) }
  if f := KToF( 273.15 ); f !=  32 { t.Errorf( "KToF( 273.15 ) != 32, result %g", f ) }
  if f := KToF( 373.15 ); f != 212 { t.Errorf( "KToF( 373.15 ) != 212, result %g", f ) }
}
