package main
import (
   "encoding/json"
   "fmt"
)
func main() {
   var y = []byte(`{"Sunday": 10}`)
   // example 1
   var m map[string]int
   json.Unmarshal(y, &m)
   fmt.Println(m)
   // example 2
   var o struct {Sunday int}
   json.Unmarshal(y, &o)
   fmt.Printf("%+v\n", o)
}
