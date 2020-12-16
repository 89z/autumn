package main

import (
   "bytes"
   "fmt"
   "io"
   "os"
   "time"
)

func GetContents(s string) (bytes.Buffer, error) {
   buf_o := bytes.Buffer{}
   open_o, e := os.Open(s)
   if e != nil {
      return buf_o, e
   }
   n, e := io.Copy(&buf_o, open_o)
   if e != nil {
      return buf_o, fmt.Errorf("%v %v", n, e)
   }
   return buf_o, nil
}

func main() {
   GetContents("100mb.file")
   time.Sleep(time.Duration(time.Minute))
}
