package main

import (
   "fmt"
   "github.com/pelletier/go-toml"
   "log"
)

func load(s string) (map[string]interface{}, error) {
   t, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return t.ToMap(), nil
}

func main() {
   m, e := load("b.toml")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
