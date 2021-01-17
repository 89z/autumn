package main

import (
   "fmt"
   "log"
)

type Map map[string]interface{}

func check(e error) {
   if e != nil {
      log.Fatal(e)
   }
}

func main() {
   m1, e := TomlGetByte([]byte("month=12\nday=31"))
   check(e)
   m2, e := TomlGetFile("manifest.toml")
   check(e)
   m3, e := TomlGetString("month=12\nday=31")
   fmt.Println(m1, m2, m3)
}
