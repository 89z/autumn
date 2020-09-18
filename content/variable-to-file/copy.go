package main

import (
   "io"
   "os"
   "strings"
)

func main() {
   oA := strings.NewReader("May\n")
   oB, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   io.Copy(oB, oA)
}
