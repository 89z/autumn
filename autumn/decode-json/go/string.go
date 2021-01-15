package main

import (
   "encoding/json"
   "fmt"
   "log"
)

func JsonFromString(s string) (map[string]interface{}, error) {
   y := []byte(s)
   m := map[string]interface{}{}
   return m, json.Unmarshal(y, &m)
}

func main() {
   s := `{"name":"GitHub","prefer_related_applications":true}`
   m, e := JsonFromString(s)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
