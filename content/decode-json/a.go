package main
import (
   "encoding/json"
   "io/ioutil"
   "log"
)
func main() {
   a, e := ioutil.ReadFile("a.json")
   if e != nil {
      log.Fatal(e)
   }
   // example 1
   var m map[string]int
   json.Unmarshal(a, &m)
   log.Print(m)
   // example 2
   var o struct{Sunday, Monday int}
   json.Unmarshal(a, &o)
   log.Printf("%+v", o)
}
