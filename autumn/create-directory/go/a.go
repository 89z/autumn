package main

import (
   "log"
   "os"
)

func mkdir(path string) error {
   err := os.Mkdir(path, os.ModeDir)
   if err != nil {
      dir, err1 := os.Stat(path)
      if err1 == nil && dir.IsDir() {
         return nil
      }
      return err
   }
   return nil
}

func main() {
   e := mkdir("March")
   if e != nil {
      log.Fatal(e)
   }
}
