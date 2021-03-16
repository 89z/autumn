package main

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
)

func main() {
   var (
      m = map[string]int{"SNG_ID": 75498415}
      b = new(bytes.Buffer)
      r = new(http.Request)
   )
   json.NewEncoder(b).Encode(m)
   r.URL = &url.URL{Scheme: "http", Host: "www.deezer.com"}
   r.Method = "POST"
   r.Body = io.NopCloser(b)
   http.DefaultClient.Do(r)
}
