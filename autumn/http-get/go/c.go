package main

import (
   "fmt"
   "net/http"
   "os"
)

func netrc() (string, string, error) {
   home, e := os.UserHomeDir()
   if e != nil { return "", "", e }
   f, e := os.Open(home + "/_netrc")
   if e != nil { return "", "", e }
   defer f.Close()
   var login, pass string
   fmt.Fscanf(f, "default login %v password %v", &login, &pass)
   return login, pass, nil
}

func main() {
   login, pass, e := netrc()
   if e != nil {
      panic(e)
   }
   req, e := http.NewRequest("GET", "https://api.github.com/rate_limit", nil)
   if e != nil {
      panic(e)
   }
   req.SetBasicAuth(login, pass)
   res, e := new(http.Client).Do(req)
   if e != nil {
      panic(e)
   }
   defer res.Body.Close()
   fmt.Println(res)
}
