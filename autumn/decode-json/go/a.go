package main

import (
   "log"
   "time"
)

type Slice []interface{}
type Map map[string]interface{}

func (m Map) M(s string) Map {
   return m[s].(map[string]interface{})
}

func (m Map) A(s string) Slice {
   return m[s].([]interface{})
}

func main() {
   m, e := JsonDecode("youtube.json")
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("videoDetails").A("keywords")[0]
   println(s == "Blue Wednesday")
   time.Sleep(time.Duration(time.Minute))
}
