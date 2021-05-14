package main

import (
   "encoding/json"
   "io"
   "net/http"
)

func main() {
   m := map[string]int{"SNG_ID": 75498415}
   r, w := io.Pipe()
   go func() {
      json.NewEncoder(w).Encode(m)
      w.Close()
   }()
   http.Post("http://www.deezer.com", "application/json", r)
}
