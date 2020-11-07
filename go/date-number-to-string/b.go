package main
import "time"

func main() {
   n := int64(1577858399)
   o := time.Unix(n, 0)
   s := o.String()
   println(s == "2019-12-31 23:59:59 -0600 CST")
}
