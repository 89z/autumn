package main

import (
   "fmt"
   "time"
)

func main() {
   today := time.Now()
   tomorrow := today.Add(24 * time.Hour)
   fmt.Println(tomorrow)
}
