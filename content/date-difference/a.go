package main
import (
   "fmt"
   "time"
)
func main() {
   o1 := time.Date(2019, time.December, 31, 0, 0, 0, 0, time.UTC)
   o2 := time.Now()
   o3 := o2.Sub(o1)
   fmt.Println(o3)
}
