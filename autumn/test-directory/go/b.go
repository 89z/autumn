package main

import (
   "log"
   "os"
)

func isDir(name string) bool {
   fi, err := os.Stat(name)
   return err == nil && fi.IsDir()
}

func main() {
   b := isDir(`C:\Users`)
   println(b)
}
