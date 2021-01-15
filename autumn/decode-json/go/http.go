package main

import (
   "encoding/json"
   "fmt"
   "log"
   "net/http"
)

func JsonFromHttp(s string) (map[string]interface{}, error) {
   o, e := http.Get(s)
   if e != nil {
      return nil, e
   }
   m := map[string]interface{}{}
   return m, json.NewDecoder(o.Body).Decode(&m)
}

func main() {
   m, e := JsonFromHttp("https://github.com/manifest.json")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
