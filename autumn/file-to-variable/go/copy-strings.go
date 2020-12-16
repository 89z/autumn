package main

import (
   "fmt"
   "io"
   "os"
   "strings"
   "time"
)

func GetContents(s string) (strings.Builder, error) {
   build_o := strings.Builder{}
   open_o, e := os.Open(s)
   if e != nil {
      return build_o, e
   }
   n, e := io.Copy(&build_o, open_o)
   if e != nil {
      return build_o, fmt.Errorf("%v %v", n, e)
   }
   return build_o, nil
}

func main() {
   GetContents("100mb.file")
   time.Sleep(time.Duration(time.Minute))
}
