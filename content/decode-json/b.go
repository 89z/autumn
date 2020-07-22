package main
import (
   "encoding/json"
   "fmt"
   "log"
   "os"
)
func main() {
   var o, e = os.Open("a.json")
   if e != nil {
      log.Fatal(e)
   }
   var m map[string]int
   json.NewDecoder(o).Decode(&m)
   fmt.Println(m)
}
