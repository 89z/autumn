package main

import (
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
   "os"
)

type Map map[string]interface{}

func TomlPutFile(source Map, dest string) error {
   y, e := toml.Marshal(source)
   if e != nil {
      return e
   }
   return ioutil.WriteFile(dest, y, os.ModePerm)
}

func main() {
   e := TomlPutFile(Map{"month": 12, "day": 31}, "a.toml")
   if e != nil {
      log.Fatal(e)
   }
}
