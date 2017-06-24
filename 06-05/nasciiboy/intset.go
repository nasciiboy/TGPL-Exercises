// Package intset provides a set of integers based on a bit vector.
package intset

import (
  "bytes"
  "fmt"
)

const wordLen = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
  words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
  word, bit := x/wordLen, uint(x%wordLen)
  return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
  word, bit := x/wordLen, uint(x%wordLen)
  for word >= len(s.words) {
    s.words = append(s.words, 0)
  }
  s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll( v ...int){
  for i := range v { s.Add( v[ i ] ) }
}

// remove x from the set
func (s *IntSet) Remove(x int) {
  word, bit := x/wordLen, uint(x%wordLen)

  if word < len(s.words) {
    s.words[word] &^= 1 << bit
  }
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] |= tword
    } else {
      s.words = append(s.words, tword)
    }
  }
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] &= tword
    } else {
      s.words = append(s.words, 0)
    }
  }
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] &^= tword
    } else {
      s.words = append(s.words, 0)
    }
  }
}

// SymmetricDifference sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] ^= tword
    } else {
      s.words = append(s.words, tword)
    }
  }
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
  var buf bytes.Buffer
  buf.WriteByte('{')
  for i, word := range s.words {
    if word == 0 {
      continue
    }
    for j := 0; j < wordLen; j++ {
      if word&(1<<uint(j)) != 0 {
        if buf.Len() > len("{") {
          buf.WriteByte(' ')
        }
        fmt.Fprintf(&buf, "%d", wordLen*i+j)
      }
    }
  }
  buf.WriteByte('}')
  return buf.String()
}

// return the number of elements
func (s *IntSet) Len() int {
  if len(s.words) == 0 { return 0 }

  return len(s.words) * wordLen
}

// remove all elements from the set
func (s *IntSet) Clear() {
  s.words = []uint{}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
  var cpy IntSet
  cpy.words = make( []uint, len(s.words) )
  copy( cpy.words, s.words )

  return &cpy
}

func (s *IntSet) Elems() []uint {
  words := make( []uint, len(s.words) )
  copy( words, s.words )

  return words
}
