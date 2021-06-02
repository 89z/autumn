package main

import (
   "io"
   "os"
)

func read(s string) ([]byte, error) {
   f, e := os.Open(s)
   if e != nil { return nil, e }
   defer f.Close()
   return io.ReadAll(f)
}

func main() {
   b, e := read("a.go")
   if e != nil {
      panic(e)
   }
   os.Stdout.Write(b)
}
