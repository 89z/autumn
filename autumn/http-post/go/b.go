package main

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func main() {
   m, b := map[string]int{"SNG_ID": 75498415}, new(bytes.Buffer)
   json.NewEncoder(b).Encode(m)
   http.Post("http://www.deezer.com", "application/json", b)
}
