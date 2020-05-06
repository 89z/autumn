package main
import (
   "fmt"
   "time"
)
func main() {
   o1 := time.Now()
   y1, _ := o1.MarshalText()
   fmt.Printf("%s\n", y1)
}
