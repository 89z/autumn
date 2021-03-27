package main
import "encoding/json"

func main() {
   var (
      b = []byte(`{"month": 12, "day": 31}`)
      m struct { Month, Day int }
   )
   json.Unmarshal(b, &m)
   println(m.Day == 31)
}
