// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
  "fmt"
  "flag"
  "crypto/sha256"
  "crypto/sha512"
)

var s384 = flag.Bool( "384", false, "prints sha384" )
var s512 = flag.Bool( "512", false, "prints sha512" )

func main() {
  flag.Parse()

  for _, arg := range flag.Args() {
    fmt.Printf( "%q\nsha256 ==> %x\n",   arg, sha256.Sum256([]byte(arg)) )
    if *s384 { fmt.Printf( "sha384 ==> %x\n", sha512.Sum384([]byte(arg)) ) }
    if *s512 { fmt.Printf( "sha512 ==> %x\n", sha512.Sum512([]byte(arg)) ) }
  }
}

//!-
// go run 04-02/nasciiboy/main.go -384 -512 "the hash program"
// "the hash program"
// sha256 ==> d89067176bff80883a064553931fe7cc086ca8e265d8bf991113f92f2e9c7baa
// sha384 ==> 15b9f3d5bd395309f21f8e144dc88650754a9d71351ddabc15c5c4228436a8bc3e3556adfbefd7ede301f36785cecb03
// sha512 ==> 7c9bfd4aaa40aa39c4f8c28a0dc20bfdec985a5e2291688c3ef8d542f66f7741bf994ac430c2effdd10a23559465c43ef0851721f4eb0716fef299d65fba8e69
