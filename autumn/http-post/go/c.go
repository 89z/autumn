package main

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

func main() {
   var r http.Request
   r.URL = &url.URL{Scheme: "http", Host: "www.deezer.com"}
   r.Method = "POST"
   r.Body = io.NopCloser(strings.NewReader(`{"sng_id": "75498415"}`))
   http.DefaultClient.Do(&r)
}
