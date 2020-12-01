package main
import "encoding/json"

func main() {
   m := Map{}
   y := []byte(in_s)
   json.Unmarshal(y, &m)
   out_s := m.M("U2").A("Boy")[0]
   println(out_s == "Twilight")
}
