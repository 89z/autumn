package main

import (
   "fmt"
   "regexp"
)

func find(pat string, sub []byte) ([]byte, error) {
   r, e := regexp.Compile(pat)
   if e != nil { return nil, e }
   b := r.Find(sub)
   if b == nil {
      return nil, fmt.Errorf("Find %v", pat)
   }
   return b, nil
}

func main() {
   b, e := find("o..", []byte("south north"))
   if e != nil {
      panic(e)
   }
   fmt.Printf("%c\n", b) // [o u t]
}
