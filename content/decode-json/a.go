package main
import (
   "encoding/json"
   "fmt"
)
func main() {
   a1 := []byte(`{"sunday": 10, "monday": 11}`)
   // example 1
   var m1 map[string]int
   json.Unmarshal(a1, &m1)
   // example 2
   var m2 struct{Sun, Mon int}
   json.Unmarshal(a1, &m2)
   // print
   fmt.Printf("%+v\n", m1)
   fmt.Printf("%+v\n", m2)
}
