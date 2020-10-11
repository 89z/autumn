package main

import (
   "bytes"
   "io/ioutil"
   "log"
   "strings"
   "testing"
)

func BenchmarkBytes(b *testing.B) {
   y, e := ioutil.ReadFile("months.txt")
   if e != nil {
      log.Fatal(e)
   }
   for n := 0; n < b.N; n++ {
      y2 := []byte("December")
      bytes.Contains(y, y2)
   }
}

func BenchmarkStrings(b *testing.B) {
   y, e := ioutil.ReadFile("months.txt")
   if e != nil {
      log.Fatal(e)
   }
   for n := 0; n < b.N; n++ {
      s := string(y)
      strings.Contains(s, "December")
   }
}
