package main

import (
   "fmt"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
)

func TomlDecode(y []byte) (map[string]interface{}, error) {
   o, e := toml.LoadBytes(y)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func main() {
   y, e := ioutil.ReadFile("manifest.toml")
   if e != nil {
      log.Fatal(e)
   }
   m, e := TomlDecode(y)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
