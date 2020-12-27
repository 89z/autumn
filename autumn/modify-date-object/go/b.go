package main

import (
   "fmt"
   "time"
)

func main() {
   now := time.Now()
   then := now.AddDate(-1, 0, 0)
   fmt.Println(then)
}
