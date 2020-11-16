package main
import "fmt"

func NumberFormat(n float32) string {
   a := []string{"", " k", " M", " G"}
   n2 := 0
   for n > 1024 {
      n /= 1024
      n2++
   }
   return fmt.Sprintf("%.3f", n) + a[n2]
}

func main() {
   s := NumberFormat(1264)
   println(s == "1.234 k")
}
