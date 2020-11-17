package main
import "fmt"

func NumberFormat(n float64) string {
   n2 := 0
   for n > 1000 {
      n /= 1000
      n2++
   }
   return fmt.Sprintf("%.3f", n) + []string{"", " k", " M", " G"}[n2]
}

func main() {
   s := NumberFormat(9012345678)
   println(s == "9.012 G")
}
