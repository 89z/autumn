package main

import (
   "fmt"
   "github.com/pelletier/go-toml"
   "log"
)

func TomlDecode(s string) (map[string]interface{}, error) {
   o, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func main() {
   m, e := TomlDecode("manifest.toml")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
