package main
import "github.com/pelletier/go-toml"

func TomlGetFile(s string) (Map, error) {
   o, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}
