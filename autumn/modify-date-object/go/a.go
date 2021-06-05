package main

import (
   "fmt"
   "time"
)

func main() {
   t := time.Now()
   t = t.Add(24 * time.Hour)
   fmt.Println(t)
}
