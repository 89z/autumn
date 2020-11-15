package main
import "fmt"

func HumanReadable(in_n float32) string {
   a := []string{"", " k", " M", " G"}
   out_n := 0
   for in_n > 1024 {
      in_n /= 1024
      out_n++
   }
   return fmt.Sprintf("%.3f", in_n) + a[out_n]
}

func main() {
   s := HumanReadable(1264)
   println(s == "1.234 k")
}
