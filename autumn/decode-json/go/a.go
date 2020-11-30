package main

import (
   "encoding/json"
   "strings"
)

// example 1
func f1(s string) map[string]int {
   o := strings.NewReader(s)
   m := map[string]int{}
   json.NewDecoder(o).Decode(&m)
   return m
}

// example 2
func f2(s string) struct{Month, Day int} {
   in_o := strings.NewReader(s)
   out_o := struct{Month, Day int}{}
   json.NewDecoder(in_o).Decode(&out_o)
   return out_o
}

// print
func main() {
   s := `{"month": 12, "day": 31}`
   m, o := f1(s), f2(s)
   println(m["day"] == 31, o.Day == 31)
}
