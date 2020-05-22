package main
import (
   "bytes"
   "encoding/json"
   "log"
)
func main() {
   a1 := []int{10, 11}
   y1, e := json.Marshal(a1)
   if e != nil {
      log.Fatal(e)
   }
   y2 := []byte("[10,11]")
   b1 := bytes.Equal(y1, y2)
   println(b1)
}
