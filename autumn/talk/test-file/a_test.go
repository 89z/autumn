package main

import (
   "os"
   "testing"
)

func BenchmarkStat(b *testing.B) {
   for n := 0; n < b.N; n++ {
      o, e := os.Stat("a_test.go")
      if os.IsNotExist(e) {
         continue
      }
      o.Mode().IsRegular()
   }
}

func BenchmarkOpen(b *testing.B) {
   for n := 0; n < b.N; n++ {
      o, e := os.Open("a_test.go")
      if os.IsNotExist(e) {
         continue
      }
      o.Close()
   }
}
