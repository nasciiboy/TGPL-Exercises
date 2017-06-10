package massconv

import "testing"

func TestPToK( t *testing.T ){
  if p := PToK( 1 ); p != 0.453592 { t.Errorf( "PToK( 1 ) != 0.453592, result %g", p ) }
  if p := PToK( 3 ); p != 1.360776 { t.Errorf( "PToK( 3 ) != 1.360776, result %g", p ) }
}

func TestKToP( t *testing.T ){
  if k := KToP( 2 ); k != 4.409246 { t.Errorf( "KToP( 2 ) != 4.409246, result %g", k ) }
  if k := KToP( 1 ); k != 2.204623 { t.Errorf( "KToP( 1 ) != 2.204623, result %g", k ) }
}
