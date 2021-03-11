package main

import (
   "log"
   "net/http"
   "strings"
)

func main() {
   b := strings.NewReader(`{"sng_id": "75498415"}`)
   r, e := http.NewRequest("POST", "http://www.deezer.com", b)
   if e != nil {
      log.Fatal(e)
   }
   http.DefaultClient.Do(r)
}
