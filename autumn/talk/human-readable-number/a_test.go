package main

import (
   "fmt"
   "math"
   "testing"
)

func For(n float64) string {
   n2 := 0
   for n > 1000 {
      n /= 1000
      n2++
   }
   return fmt.Sprintf("%.3f", n) + []string{"", " k", " M", " G"}[n2]
}

func BenchmarkFor(b *testing.B) {
   for n := 0; n < b.N; n++ {
      For(5678901234)
   }
}

func Log(n float64) string {
   n2 := int(math.Log(n) / math.Log(1000))
   n /= math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n) + []string{"", " k", " M", " G"}[n2]
}

func BenchmarkLog(b *testing.B) {
   for n := 0; n < b.N; n++ {
      Log(5678901234)
   }
}

func Log10(n float64) string {
   n2 := int(math.Log10(n) / 3)
   n /= math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n) + []string{"", " k", " M", " G"}[n2]
}

func BenchmarkLog10(b *testing.B) {
   for n := 0; n < b.N; n++ {
      Log10(5678901234)
   }
}
