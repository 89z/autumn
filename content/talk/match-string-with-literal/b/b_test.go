package main

import (
   "strings"
   "testing"
)

func hasPrefixByte(s string, prefix byte) bool {
   return s[0] == prefix
}
func BenchmarkIndexByte(o *testing.B) {
   for n := 0; n < o.N; n++ {
      hasPrefixByte("May", 'M')
   }
}

func BenchmarkStrings(o *testing.B) {
   for n := 0; n < o.N; n++ {
      strings.HasPrefix("May", "M")
   }
}
