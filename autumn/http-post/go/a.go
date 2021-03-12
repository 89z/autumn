package main

import (
   "bytes"
   "encoding/json"
   "log"
   "net/http"
)

func main() {
   m, b := struct{SngId int}{75498415}, new(bytes.Buffer)
   json.NewEncoder(b).Encode(m)
   r, e := http.NewRequest("POST", "http://www.deezer.com", b)
   if e != nil {
      log.Fatal(e)
   }
   http.DefaultClient.Do(r)
}
