package main
import (
   "fmt"
   "time"
)
func main() {
   n := int64(365 * 24 * 60 * 60)
   o := time.Unix(n, 0)
   fmt.Println(o)
}
