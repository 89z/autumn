package main
import "time"

func main() {
   n := int64(1577858399)
   o := time.Unix(n, 0)
   s := o.Format(time.RFC3339)
   println(s == "2019-12-31T23:59:59-06:00")
}
