package main
import (
   "encoding/json"
   "fmt"
)
func main() {
   y := []byte(`
{"Sunday": 10}
`)
   // example 1
   m := map[string]int{}
   json.Unmarshal(y, &m)
   // example 2
   o := struct{Sunday int}{}
   json.Unmarshal(y, &o)
   // print
   fmt.Printf("%v %+v\n", m, o)
}
