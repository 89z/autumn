package main

import (
   "fmt"
   "github.com/pelletier/go-toml"
   "log"
)

func load(b []byte) (map[string]interface{}, error) {
   t, e := toml.LoadBytes(b)
   if e != nil {
      return nil, e
   }
   return t.ToMap(), nil
}

func main() {
   m, e := load([]byte("month=12\nday=31"))
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
