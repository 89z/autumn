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
   m1, e := JsonGetFile("manifest.json")
   check(e)
   m2, e := JsonGetHttp("https://github.com/manifest.json")
   check(e)
   m3, e := JsonGetString(`{"month": 12, "day": 31}`)
   check(e)
   fmt.Println(m1, m2, m3)
}
