package main

import (
   "os"
   "strings"
)

func main() {
   oA := strings.NewReader("May\n")
   oB, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   oA.WriteTo(oB)
}
