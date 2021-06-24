package main

import (
   "io"
   "os"
)

func read(s string) ([]byte, error) {
   f, err := os.Open(s)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   return io.ReadAll(f)
}

func main() {
   b, err := read("a.go")
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(b)
}
