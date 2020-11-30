package main
import "encoding/json"

// example 1
func f1(s string) map[string]int {
   y := []byte(s)
   m := map[string]int{}
   json.Unmarshal(y, &m)
   return m
}

// example 2
func f2(s string) struct{Month, Day int} {
   y := []byte(s)
   o := struct{Month, Day int}{}
   json.Unmarshal(y, &o)
   return o
}

// print
func main() {
   s := `{"month": 12, "day": 31}`
   m, o := f1(s), f2(s)
   println(m["day"] == 31, o.Day == 31)
}
