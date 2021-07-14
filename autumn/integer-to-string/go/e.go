package main

import (
   "text/template"
   "strings"
)

func format(s string, v interface{}) string {
   t, err := new(template.Template).Parse(s)
   if err != nil {
      return ""
   }
   b := new(strings.Builder)
   t.Execute(b, v)
   return b.String()
}

func main() {
   s := struct{N int}{100}
   f := format("{{.N}}", s)
   println(f == "100")
}
