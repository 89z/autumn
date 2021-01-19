package main
import "encoding/json"

func JsonLoad(content string) (Map, error) {
   data := []byte(s)
   v := Map{}
   return v, json.Unmarshal(data, &v)
}
