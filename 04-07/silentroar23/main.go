package main

import (
	"fmt"
	"unicode/utf8"
)

func rev(r []byte) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

// Reverse reverse a []byte slice that represents an UTF-8 encoded string in place.
// the idea: a rune takes positions from i to j before reverse will take positions from len(s)-j-1 to
// len(s)-i-1. So reverse the whole slice will put the runes in the right place. Reverse inside every
// rune will put each byte in the rune in the right place
// an example:
// 1,2,3 | 4,5 | 6
// 3,2,1 | 5,4 | 6
// 6 | 4,5 | 1,2,3
func Reverse(s []byte) {
	for i, width := 0, 0; i < len(s); i += width {
		_, width = utf8.DecodeRune(s[i:])
		rev(s[i : i+width])
	}
	rev(s)
}

func main() {
	data := [][]byte{
		[]byte("start"),
		[]byte("tert"),
		[]byte("12Ø34"),
		[]byte("123Ø4"),
		[]byte("1Ø234"),
		[]byte("¤2Ø34"),
		[]byte("1¤3Ø4"),
		[]byte("1Ø2¤4"),
		[]byte("Øe¤¥næn"),
	}

	for i := range data {
		fmt.Printf("%s | ", data[i])
		Reverse(data[i])
		fmt.Printf("%s | ", data[i])
		Reverse(data[i])
		fmt.Printf("%s\n", data[i])
	}
}
