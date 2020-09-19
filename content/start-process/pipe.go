package main

import (
   "fmt"
   "time"
)

func main() {
   fmt.Println("end")
   o := time.Duration(5 * time.Second)
   time.Sleep(o)
}
