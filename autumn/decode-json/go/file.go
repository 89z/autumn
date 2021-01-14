package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "log"
)

func JsonDecode(s string) (map[string]interface{}, error) {
   y, e := ioutil.ReadFile(s)
   if e != nil {
      return nil, e
   }
   m := map[string]interface{}{}
   return m, json.Unmarshal(y, &m)
}

func main() {
   m, e := JsonDecode("manifest.json")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}

