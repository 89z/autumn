package main

import (
   "fmt"
   "regexp"
)

func find(pat string, sub []byte) ([]byte, error) {
   re, err := regexp.Compile(pat)
   if err != nil { return nil, err }
   match := re.Find(sub)
   if match == nil {
      return nil, fmt.Errorf("Find %v", pat)
   }
   return match, nil
}

func main() {
   match, err := find("o..", []byte("south north"))
   if err != nil {
      panic(err)
   }
   fmt.Printf("%c\n", match) // [o u t]
}
