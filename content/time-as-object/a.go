package main
import (
   "fmt"
   "time"
)
func main() {
   // example 1
   t1 := time.Now()
   // example 2
   t2 := time.Unix(86400, 0)
   // print
   fmt.Print(t1, "\n", t2, "\n")
}
