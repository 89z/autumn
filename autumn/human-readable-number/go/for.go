package main
import "fmt"

func numberFormat(d float64) string {
   e, f := d, 0
   for e >= 1e3 {
      e /= 1e3
      f++
   }
   return fmt.Sprintf("%.3f", e) + []string{"", " k", " M", " G"}[f]
}

func main() {
   s := numberFormat(9012345678)
   println(s == "9.012 G")
}
