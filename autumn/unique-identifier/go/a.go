package main

import (
   "fmt"
   "log"
   "math/big"
   "time"
)

func idDecode(s string, year int) (time.Time, error) {
   n, ok := new(big.Int).SetString(s, 36)
   if ! ok {
      return time.Time{}, fmt.Errorf("%v", s)
   }
   return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Add(
      time.Duration(n.Int64()) * time.Second,
   ), nil
}

func main() {
   t, e := idDecode("6dv3d", 2020)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(t)
}
