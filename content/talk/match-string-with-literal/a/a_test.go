package main

import (
   "bytes"
   "strings"
   "testing"
)

const s = `January
February
March
April
May
June
July
August
September
October
November
December`

func BenchmarkBytes(b *testing.B) {
   y := []byte(s)
   for n := 0; n < b.N; n++ {
      y2 := []byte("December")
      bytes.Contains(y, y2)
   }
}

func BenchmarkStrings(b *testing.B) {
   y := []byte(s)
   for n := 0; n < b.N; n++ {
      s2 := string(y)
      strings.Contains(s2, "December")
   }
}
