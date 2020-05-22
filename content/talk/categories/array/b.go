package main
import (
   "bytes"
   "encoding/json"
)
func main() {
   a1 := []int{10, 11}
   y1, e1 := json.Marshal(a1)
   if e1 == nil {
      y2 := []byte("[10,11]")
      b1 := bytes.Equal(y1, y2)
      println(b1)
   }
}
