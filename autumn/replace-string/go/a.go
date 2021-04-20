package main

import (
   "net/url"
   "regexp"
   "strings"
)

func main() {
   s := "Have You Ever Really Loved A Woman?"
   {
      t := url.QueryEscape(s)
      println(t == "Have+You+Ever+Really+Loved+A+Woman%3F")
   }
   {
      t := s
      for _, r := range `"*/:<>?\|` {
         t = strings.ReplaceAll(t, string(r), "+")
      }
      println(t == "Have You Ever Really Loved A Woman+")
   }
   {
      var a []string
      for _, r := range `"*/:<>?\|` {
         a = append(a, string(r), "+")
      }
      t := strings.NewReplacer(a...).Replace(s)
      println(t == "Have You Ever Really Loved A Woman+")
   }
   {
      t := regexp.MustCompile(`["*/:<>?\|]`).ReplaceAllString(s, "+")
      println(t == "Have You Ever Really Loved A Woman+")
   }
}
