package main
import (
   "encoding/json"
   "fmt"
   "os"
)
func main() {
   o, e := os.Open("a.json")
   if e != nil {
      os.Exit(1)
   }
   m := map[string]int{}
   json.NewDecoder(o).Decode(&m)
   fmt.Println(m)
}
