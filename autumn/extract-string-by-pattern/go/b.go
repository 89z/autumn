package main

import (
   "fmt"
   "regexp"
)

func findString(pat, sub string) (string, error) {
   re, err := regexp.Compile(pat)
   if err != nil { return "", err }
   match := re.FindString(sub)
   if match == "" {
      return "", fmt.Errorf("FindString %v", pat)
   }
   return match, nil
}

func main() {
   match, err := findString("o..", "south north")
   if err != nil {
      panic(err)
   }
   println(match) // out
}
