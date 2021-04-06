package main
import "strconv"

func main() {
   n, e := strconv.Atoi("100")
   if e != nil {
      panic(e)
   }
   println(n == 100)
}
