package main

import (
   "os"
   "strings"
)

func main() {
   o := strings.NewReader("May\n")
   o.WriteTo(os.Stdout)
}
