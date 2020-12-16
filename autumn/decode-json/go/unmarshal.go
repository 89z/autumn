package main
import "encoding/json"

func JsonDecode(s string) (Map, error) {
   y := []byte(s)
   m := Map{}
   e := json.Unmarshal(y, &m)
   if e != nil {
      return nil, e
   }
   return m, nil
}
