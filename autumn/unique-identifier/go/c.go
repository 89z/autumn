package main
import "strconv"

func encode36(n int) string {
   return strconv.FormatInt(
      int64(n), 36,
   )
}

func main() {
   n := 1609480799
   s := encode36(n)
   println(s == "qm8rbz")
}
