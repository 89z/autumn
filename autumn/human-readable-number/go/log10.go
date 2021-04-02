package main

import (
   "fmt"
   "math"
)

func numberFormat(d float64) string {
   e := int(math.Log10(d)) / 3
   f := d / math.Pow10(e * 3)
   return fmt.Sprintf("%.3f", f) + []string{"", " k", " M", " G"}[e]
}

func main() {
   s := numberFormat(9012345678)
   println(s == "9.012 G")
}
