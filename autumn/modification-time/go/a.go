package main

import (
   "os"
   "time"
)

func main() {
   o := time.Unix(400_000_000, 0)
   os.Chtimes("a.go", o, o)
}
