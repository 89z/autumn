package main
import "time"
func main() {
   n := int64(366 * 24 * 60 * 60)
   o := time.Unix(n, 0)
   s := o.Format("Mon Jan 2 2006")
   println(s == "Fri Jan 1 1971")
}
