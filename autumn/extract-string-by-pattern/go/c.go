package main

import (
   "fmt"
   "regexp"
)

func findStringSubmatch(pat, sub string) ([]string, error) {
   re, err := regexp.Compile(pat)
   if err != nil { return nil, err }
   match := re.FindStringSubmatch(sub)
   if match == nil {
      return nil, fmt.Errorf("FindStringSubmatch %v", pat)
   }
   return match, nil
}

func main() {
   match, err := findStringSubmatch("o(..)", "south north")
   if err != nil {
      panic(err)
   }
   fmt.Println(match) // [out, ut]
}
