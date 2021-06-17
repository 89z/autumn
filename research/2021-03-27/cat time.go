package main

import (
   "fmt"
   "time"
)

func main() {
   t := time.Unix(0x5555_5555, 0)
   // 2015-05-14 21:09:25 -0500 CDT
   fmt.Println(t)
   // 2015-05-15 02:09:25 +0000 UTC
   fmt.Println(t.UTC())
}
