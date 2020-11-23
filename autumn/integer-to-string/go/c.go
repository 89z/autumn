package main
import "strconv"

func main() {
   n := int64(100)
   s := strconv.FormatInt(n, 10)
   println(s == "100")
}
