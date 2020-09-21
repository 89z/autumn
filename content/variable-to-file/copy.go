package main

import (
   "io"
   "os"
   "strings"
)

func main() {
   o := strings.NewReader("May\n")
   o1, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   io.Copy(o1, o)
}
