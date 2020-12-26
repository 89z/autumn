package main

import (
   "fmt"
   "time"
)

func main() {
   o := time.Now().AddDate(0, 0, -1)
   fmt.Println(o)
}
