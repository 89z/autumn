package main

import (
   "fmt"
   "net/url"
)

func main() {
   for _, s := range []string{
      "https://golang.org/pkg", // "https" "golang.org" "/pkg"
      "//golang.org",           // ""      "golang.org" ""
      "pkg",                    // ""      ""           "pkg"
   } {
      p, e := url.Parse(s)
      if e != nil {
         panic(e)
      }
      fmt.Printf("%q %q %q\n", p.Scheme, p.Host, p.Path)
   }
}
