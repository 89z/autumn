package main

import (
   "fmt"
   "log"
)

type Map map[string]interface{}

func main() {
   m1, e := TomlGetByte([]byte("month=12\nday=31"))
   if e != nil {
      log.Fatal(e)
   }
   m2, e := TomlGetFile("manifest.toml")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m1, m2)
}
