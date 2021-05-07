package main

import (
   "fmt"
   "regexp"
)

func findSubmatch(pat string, sub []byte) ([][]byte, error) {
   re, err := regexp.Compile(pat)
   if err != nil { return nil, err }
   match := re.FindSubmatch(sub)
   if match == nil {
      return nil, fmt.Errorf("FindSubmatch %v", pat)
   }
   return match, nil
}

func main() {
   match, err := findSubmatch("o(..)", []byte("south north"))
   if err != nil {
      panic(err)
   }
   fmt.Printf("%c\n", match) // [[o u t] [u t]]
}
