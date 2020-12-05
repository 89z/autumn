package main

import (
   "fmt"
   "log"
   "time"
)

func main() {
   o, e := time.Parse(time.RFC3339[:10], "2019-12-31")
   if e != nil {
      log.Fatal(e)
   }
   n := time.Now().Sub(o).Hours()
   fmt.Println(n)
}
