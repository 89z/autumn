package main
import "github.com/pelletier/go-toml"

func TomlGetByte(y []byte) (Map, error) {
   o, e := toml.LoadBytes(y)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}
