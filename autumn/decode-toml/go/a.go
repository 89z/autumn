package main

import (
   "encoding/json"
   "github.com/pelletier/go-toml"
   "os"
)

const s = `{
   "This is a": "TOML document",
   "name": "Tom Preston-Werner",
   "title": "TOML Example"
}`

func main() {
   y := []byte(s)
   m := map[string]string{}
   json.Unmarshal(y, &m)
   toml.NewEncoder(os.Stdout).QuoteMapKeys(true).Encode(m)
}
