package main

import (
   "io/ioutil"
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

func GetContents(s string) (string, error) {
   y, e := ioutil.ReadFile(s)
   if e != nil {
      return "", e
   }
   return string(y), nil
}

func main() {
   get_s, e := GetContents("youtube.json")
   if e != nil {
      log.Fatal(e)
   }
   m, e := JsonDecode(get_s)
   if e != nil {
      log.Fatal(e)
   }
   key_s := m.M("videoDetails").A("keywords")[0]
   println(key_s == "Blue Wednesday")
   time.Sleep(time.Duration(time.Minute))
}
