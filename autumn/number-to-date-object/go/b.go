package main

import (
   "fmt"
   "time"
)

func main() {
   n := 0x5555_5555
   t := time.Unix(int64(n), 0)
   fmt.Println(t)
}
