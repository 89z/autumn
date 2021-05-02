package main
import "fmt"

func numberFormat(d float64) string {
   var e int
   for d >= 1000 {
      d /= 1000
      e++
   }
   return fmt.Sprintf("%.3f", d) + []string{"", " k", " M", " G"}[e]
}

func main() {
   s := numberFormat(9012345678)
   println(s == "9.012 G")
}
