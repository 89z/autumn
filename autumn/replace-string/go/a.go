package main

import (
   "net/url"
   "regexp"
   "strings"
)

func main() {
   s := "Have You Ever Really Loved A Woman?"
   {
      t := regexp.MustCompile(`["*/:<>?\|]`).ReplaceAllString(s, "+")
      println(t == "Have You Ever Really Loved A Woman+")
   }
}
