package main

import (
   "bytes"
   "strings"
   "testing"
)

func BenchmarkBytes(o *testing.B) {
   y, y2 := []byte("May"), []byte("M")
   for n := 0; n < o.N; n++ {
      bytes.HasPrefix(y, y2)
   }
}

func hasPrefixByte(s string, prefix byte) bool {
   return s[0] == prefix
}
func BenchmarkIndexByte(o *testing.B) {
   for n := 0; n < o.N; n++ {
      hasPrefixByte("May", 'M')
   }
}

func hasPrefixString(s, prefix string) bool {
   return s[:1] == prefix
}
func BenchmarkIndexString(o *testing.B) {
   for n := 0; n < o.N; n++ {
      hasPrefixString("May", "M")
   }
}

func BenchmarkStrings(o *testing.B) {
   for n := 0; n < o.N; n++ {
      strings.HasPrefix("May", "M")
   }
}
