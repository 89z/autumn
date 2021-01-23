package main

import (
   "os"
   "time"
)

func touch(filename string, sec int64) error {
   t := time.Unix(sec, 0)
   return os.Chtimes(filename, t, t)
}

func main() {
   touch("a.go", 400_000_000)
}
