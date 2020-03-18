package main
import (
   "encoding/json"
   "fmt"
)
func main() {
   a1 := []byte(`
{"Sunday": 10, "Monday": 11}
`)
   // example 1
   var m1 map[string]int
   json.Unmarshal(a1, &m1)
   // example 2
   var m2 struct{Sunday, Monday int}
   json.Unmarshal(a1, &m2)
   // print
   fmt.Printf("%+v\n", []interface{}{m1, m2})
}
