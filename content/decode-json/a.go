package main
import (
   "encoding/json"
   "fmt"
   "io/ioutil"
)
func main() {
   a1, _ := ioutil.ReadFile("a.json")
   // example 1
   var m1 map[string]int
   json.Unmarshal(a1, &m1)
   fmt.Println(m1)
   // example 2
   var o1 struct{Sunday, Monday int}
   json.Unmarshal(a1, &o1)
   fmt.Printf("%+v\n", o1)
}
