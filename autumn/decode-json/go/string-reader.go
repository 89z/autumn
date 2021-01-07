package main

import (
   "encoding/json"
   "fmt"
   "log"
   "strings"
)

func JsonDecode(s string) ([]interface{}, error) {
   o, a := strings.NewReader(s), []interface{}{}
   return a, json.NewDecoder(o).Decode(&a)
}

func main() {
   a, e := JsonDecode(`[{"month":12,"day":31},{"month":11,"day":30}]`)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
