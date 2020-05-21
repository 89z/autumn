package main
import (
   "encoding/json"
   "fmt"
)
func main() {
   a := []string{"Sun", "Mon"}
   y, e := json.Marshal(a)
   if e == nil {
      fmt.Printf("%s\n", y)
   }
}
