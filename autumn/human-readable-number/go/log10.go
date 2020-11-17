package main

import (
   "fmt"
   "math"
)

func NumberFormat(n float64) string {
   n2 := int(math.Log10(n)) / 3
   n /= math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n) + []string{"", " k", " M", " G"}[n2]
}

func main() {
   s := NumberFormat(1234)
   println(s)
}
