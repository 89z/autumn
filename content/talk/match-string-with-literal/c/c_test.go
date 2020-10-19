package main

import (
   "strings"
   "testing"
)

func BenchmarkStrings(o *testing.B) {
   for n := 0; n < o.N; n++ {
      strings.HasPrefix("June", "Ju")
   }
}

func hasPrefix(s, prefix string) bool {
   return s[:len(prefix)] == prefix
}
func BenchmarkIndexString(o *testing.B) {
   for n := 0; n < o.N; n++ {
      hasPrefix("June", "Ju")
   }
}

