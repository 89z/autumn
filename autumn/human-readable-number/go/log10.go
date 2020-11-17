package main

import (
   "fmt"
   "math"
)

func NumberFormat(n float64) string {
   n2 := int(math.Log10(n) / 3)
   n /= math.Pow(1024, float64(n2))
   return fmt.Sprintf("%.3f", n) + []string{"", " k", " M", " G"}[n2]
}

func main() {
   s := NumberFormat(1264)
   println(s)
}
