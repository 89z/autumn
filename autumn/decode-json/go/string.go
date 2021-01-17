package main
import "encoding/json"

func JsonGetString(s string) (Map, error) {
   y := []byte(s)
   m := Map{}
   return m, json.Unmarshal(y, &m)
}
