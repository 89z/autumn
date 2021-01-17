package main
import "github.com/pelletier/go-toml"

func TomlGetString(s string) (Map, error) {
   o, e := toml.Load(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}
