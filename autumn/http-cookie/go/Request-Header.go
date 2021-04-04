package main

import (
   "log"
   "net/http"
)

func main() {
   r, e := http.NewRequest("GET", "https://stackoverflow.com", nil)
   if e != nil {
      log.Fatal(e)
   }
   r.Header.Set("Cookie", "west=left")
}
