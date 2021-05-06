package main

import (
   "fmt"
   "regexp"
)

func findStringSubmatch(pat, sub string) ([]string, error) {
   r, e := regexp.Compile(pat)
   if e != nil { return nil, e }
   a := r.FindStringSubmatch(sub)
   if a == nil {
      return nil, fmt.Errorf("FindStringSubmatch %v", pat)
   }
   return a, nil
}

func main() {
   a, e := findStringSubmatch("o(..)", "south north")
   if e != nil {
      panic(e)
   }
   fmt.Println(a) // [out, ut]
}
