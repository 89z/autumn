package main
import (
   "fmt"
   "time"
)
func main() {
   var n int64 = 365 * 24 * 60 * 60
   var o = time.Unix(n, 0)
   fmt.Println(o)
}
