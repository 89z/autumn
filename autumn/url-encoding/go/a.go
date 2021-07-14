package main

import (
   "fmt"
   "net/url"
)

func main() {
   for _, s := range []string{
      "https://pkg.go.dev/std", // "https" "pkg.go.dev" "/std"
      "//pkg.go.dev",           // ""      "pkg.go.dev" ""
      "std",                    // ""      ""           "std"
   } {
      p, e := url.Parse(s)
      if e != nil {
         panic(e)
      }
      fmt.Printf("%q %q %q\n", p.Scheme, p.Host, p.Path)
   }
}
