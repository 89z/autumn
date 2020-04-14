package main
import (
   "fmt"
   "os"
)
func main() {
   // example 1
   fmt.Println("Sunday")
   // example 2
   fmt.Fprintln(os.Stderr, "Sunday")
   // example 3
   os.Stdout.WriteString("Sunday\n")
   // example 4
   println("Sunday")
   // example 5
   print("Sunday\n")
}
