package main

import (
   "fmt"
   "github.com/pelletier/go-toml"
   "log"
)

func TomlDecode(s string) (map[string]interface{}, error) {
   o, e := toml.Load(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func main() {
   s := `["excpt.h"]
package = 'Microsoft.VisualC.140.CRT.Headers.Msi'
payload = [ 'VC_CRT.Headers.msi', 'cab1.cab' ]`
   m, e := TomlDecode(s)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(m)
}
