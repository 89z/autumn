package main

import (
   "fmt"
   "net/http"
   "os"
)

func get(addr string) (*http.Response, error) {
   home, err := os.UserHomeDir()
   if err != nil { return nil, err }
   file, err := os.Open(home + "/_netrc")
   if err != nil { return nil, err }
   defer file.Close()
   var login, pass string
   fmt.Fscanf(file, "default login %v password %v", &login, &pass)
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil { return nil, err }
   req.SetBasicAuth(login, pass)
   return new(http.Client).Do(req)
}

func main() {
   res, err := get("https://api.github.com/rate_limit")
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   fmt.Println(res)
}
