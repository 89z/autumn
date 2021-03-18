package main

import (
   "bytes"
   "encoding/ascii85"
)

func main() {
   
   a := []byte{10, 11, 12, 13}
   
   
   
   // example 2
   var buf bytes.Buffer
   e := ascii85.NewEncoder(&buf)
   e.Write(a)
   e.Close()
   
   
   
   // print
   println(string(b), buf.String())
   
   
   
}
