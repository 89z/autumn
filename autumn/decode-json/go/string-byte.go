package main

import (
   "encoding/json"
   "fmt"
   "log"
)

func JsonDecode(s string) ([]interface{}, error) {
   y, a := []byte(s), []interface{}{}
   return a, json.Unmarshal(y, &a)
}

func main() {
   a, e := JsonDecode(`[{"month":12,"day":31},{"month":11,"day":30}]`)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
