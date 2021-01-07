package main

import (
   "encoding/json"
   "fmt"
   "log"
   "net/http"
)

func JsonDecode(s string) (map[string]string, error) {
   o, e := http.Get(s)
   if e != nil {
      return nil, e
   }
   m := map[string]string{}
   return m, json.NewDecoder(o.Body).Decode(&m)
}

func main() {
   m, e := JsonDecode("https://www.digitalocean.com/page-data/app-data.json")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
