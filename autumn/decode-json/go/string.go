package main

import (
   "encoding/json"
   "fmt"
   "log"
)

func JsonDecode(s string) (map[string]interface{}, error) {
   y := []byte(s)
   m := map[string]interface{}{}
   return m, json.Unmarshal(y, &m)
}

func main() {
   m, e := JsonDecode(`{"name":"GitHub","prefer_related_applications":true}`)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
