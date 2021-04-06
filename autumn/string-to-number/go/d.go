package main
import "strconv"

func main() {
   n, e := strconv.ParseFloat("99.9", 64)
   if e != nil {
      panic(e)
   }
   println(n == 99.9)
}
