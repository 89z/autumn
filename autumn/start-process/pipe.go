package main

import (
   "fmt"
   "time"
)

func main() {
   fmt.Println("END")
   o := time.Duration(5 * time.Second)
   time.Sleep(o)
}
