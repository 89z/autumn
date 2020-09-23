package main

import (
   "fmt"
   "os"
)

func main() {
   n := 9.9
   fmt.Fprintf(os.Stdout, "%.0f\n", n)
}
