package main

import (
   "log"
   "os"
)

func isFile(name string) bool {
   fi, err := os.Stat(name)
   return err == nil && ! fi.IsDir()
}

func main() {
   b := isFile("a.go")
   println(b)
}
