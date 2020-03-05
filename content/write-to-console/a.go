package main
import (
   "fmt"
   "os"
)
func main() {
   // example 1
   println("sunday")
   // example 2
   fmt.Println("sunday")
   // example 3
   fmt.Fprintln(os.Stderr, "sunday")
}
