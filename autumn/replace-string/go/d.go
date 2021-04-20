package main

import (
   "net/url"
   "regexp"
   "strings"
)

func main() {
   s := "Have You Ever Really Loved A Woman?"
   {
      t := s
      for _, r := range `"*/:<>?\|` {
         t = strings.ReplaceAll(t, string(r), "+")
      }
      println(t == "Have You Ever Really Loved A Woman+")
   }
}
