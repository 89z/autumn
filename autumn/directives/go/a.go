package main

import (
   "embed"
   "os"
)

//go:embed a.go
var content embed.FS

func main() {
   b, e := content.ReadFile("a.go")
   if e != nil {
      panic(e)
   }
   os.Stdout.Write(b)
}
