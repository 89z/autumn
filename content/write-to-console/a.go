package main
import (
   "fmt"
   "os"
)
func main() {
   // example 1
   println("Sunday")
   // example 2
   fmt.Println("Sunday")
   // example 3
   fmt.Fprintln(os.Stderr, "Sunday")
}
