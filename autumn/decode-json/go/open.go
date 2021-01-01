package main

import (
   "encoding/json"
   "fmt"
   "log"
   "os"
)

func JsonDecode(s string) ([]interface{}, error) {
   o, e := os.Open(s)
   if e != nil {
      return nil, e
   }
   defer o.Close()
   a := []interface{}{}
   return a, json.NewDecoder(o).Decode(&a)
}

func main() {
   a, e := JsonDecode("a.json")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
