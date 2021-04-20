package main

import (
   "net/url"
   "regexp"
   "strings"
)

func main() {
   s := "Have You Ever Really Loved A Woman?"
   {
      var a []string
      for _, r := range `"*/:<>?\|` {
         a = append(a, string(r), "+")
      }
      t := strings.NewReplacer(a...).Replace(s)
      println(t == "Have You Ever Really Loved A Woman+")
   }
}
