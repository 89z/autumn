package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "log"
)

func JsonDecode(s string) ([]interface{}, error) {
   y, e := ioutil.ReadFile(s)
   if e != nil {
      return nil, e
   }
   a := []interface{}{}
   return a, json.Unmarshal(y, &a)
}

func main() {
   a, e := JsonDecode("a.json")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
