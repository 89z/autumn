package main
import "fmt"

func NumberFormat(n float64) string {
   n2, n3 := n, 0
   for n2 >= 1e3 {
      n2 /= 1e3
      n3++
   }
   return fmt.Sprintf("%.3f", n2) + []string{"", " k", " M", " G"}[n3]
}

func main() {
   s := NumberFormat(9012345678)
   println(s == "9.012 G")
}
