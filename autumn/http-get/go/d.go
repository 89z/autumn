package main

import (
   "fmt"
   "net/http"
   "os"
)

func netrc(addr string) (*http.Request, error) {
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
   return req, nil
}

func main() {
   req, err := netrc("https://api.github.com/rate_limit")
   if err != nil {
      panic(err)
   }
   res, err := new(http.Client).Do(req)
   if err != nil {
      panic(err)
   }
   fmt.Println(res)
}
