package main

import (
   "net/http"
   "strings"
)

func main() {
   b := strings.NewReader(`{"sng_id": "75498415"}`)
   http.Post("http://www.deezer.com", "application/json", b)
}
