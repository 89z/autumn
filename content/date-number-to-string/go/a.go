package main
import "time"

func main() {
   n := int64(1577858399)
   o := time.Unix(n, 0)
   // example 1
   s1 := o.String()
   // example 2
   s2 := o.Format(time.RFC3339)
   // print
   println(
      s1 == "2019-12-31 23:59:59 -0600 CST",
      s2 == "2019-12-31T23:59:59-06:00",
   )
}
