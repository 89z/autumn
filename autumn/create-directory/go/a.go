package main

import (
   "log"
   "os"
)

func mkdir(path string) error {
   dir, err := os.Stat(path)
   if err == nil && dir.IsDir() {
      return nil
   }
   return os.Mkdir(path, os.ModeDir)
}

func main() {
   e := mkdir("north")
   if e != nil {
      log.Fatal(e)
   }
}
