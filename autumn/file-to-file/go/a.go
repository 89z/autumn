package main

import (
   "io"
   "log"
   "os"
)

func copyFile(source, dest string) (int64, error) {
   open, e := os.Open(source)
   if e != nil {
      return 0, e
   }
   create, e := os.Create(dest)
   if e != nil {
      return 0, e
   }
   return io.Copy(create, open)
}

func main() {
   _, e := copyFile("b.go", "c.go")
   if e != nil {
      log.Fatal(e)
   }
}
