package main

import (
   "fmt"
   "regexp"
)

func findString(pat, sub string) (string, error) {
   r, e := regexp.Compile(pat)
   if e != nil { return "", e }
   s := r.FindString(sub)
   if s == "" {
      return "", fmt.Errorf("FindString %v", pat)
   }
   return s, nil
}

func main() {
   s, e := findString("o..", "south north")
   if e != nil {
      panic(e)
   }
   println(s) // out
}
