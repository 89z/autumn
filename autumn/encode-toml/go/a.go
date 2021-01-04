package main
import "log"
type Map map[string]interface{}

func main() {
   m := Map{"month": 5, "day": 4}
   e := TomlEncode("a.toml", m)
   if e != nil {
      log.Fatal(e)
   }
}
