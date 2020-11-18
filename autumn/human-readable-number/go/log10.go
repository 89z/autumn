package main

import (
   "fmt"
   "math"
)

func NumberFormat(n float64) string {
   n2 := int(math.Log10(n)) / 3
   n3 := n / math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n3) + []string{"", " k", " M", " G"}[n2]
}

func main() {
   s := NumberFormat(9012345678)
   println(s == "9.012 G")
}
