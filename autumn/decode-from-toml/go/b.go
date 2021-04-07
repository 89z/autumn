package main

import (
   "fmt"
   "github.com/pelletier/go-toml"
)

func load(s string) (map[string]interface{}, error) {
   t, e := toml.LoadFile(s)
   if e != nil { return nil, e }
   return t.ToMap(), nil
}

func main() {
   m, e := load("a.toml")
   if e != nil {
      panic(e)
   }
   fmt.Println(m)
}
