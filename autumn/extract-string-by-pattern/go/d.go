package main

import (
   "fmt"
   "regexp"
)

func findSubmatch(pat string, sub []byte) ([][]byte, error) {
   r, e := regexp.Compile(pat)
   if e != nil { return nil, e }
   a := r.FindSubmatch(sub)
   if a == nil {
      return nil, fmt.Errorf("FindSubmatch %v", pat)
   }
   return a, nil
}

func main() {
   a, e := findSubmatch("o(..)", []byte("south north"))
   if e != nil {
      panic(e)
   }
   fmt.Printf("%c\n", a) // [[o u t] [u t]]
}
