package main

import (
   "text/template"
   "strings"
)

func format(s string, v interface{}) string {
   t, b := new(template.Template), new(strings.Builder)
   template.Must(t.Parse(s)).Execute(b, v)
   return b.String()
}

func main() {
   n := 999
   s := format("{{.}}", n)
   println(s == "999")
}
