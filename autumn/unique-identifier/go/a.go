package main

import (
   "fmt"
   "math/big"
   "time"
)

func idDecode(s string, year int) (*time.Time, error) {
   z, ok := new(big.Int).SetString(s, 36)
   if ! ok {
      return nil, fmt.Errorf("%q invalid", s)
   }
   t := time.Date(year, 1, 1, 0, 0, int(z.Int64()), 0, time.UTC)
   return &t, nil
}

func main() {
   t, err := idDecode("6dv3d", 2020)
   if err != nil {
      panic(err)
   }
   fmt.Println(t)
}
