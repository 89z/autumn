package main

import (
   "encoding/json"
   "strings"
)

func encode(s string) string {
   b := new(strings.Builder)
   enc := json.NewEncoder(b)
   enc.SetEscapeHTML(false)
   enc.Encode(s)
   return b.String()
}

func main() {
   s := encode("west & east")
   println(s == "\"west & east\"\n")
}
