package main
import (
   "fmt"
   "os"
   "time"
)
func main() {
   y, e := time.Now().MarshalText()
   if e != nil {
      os.Exit(1)
   }
   fmt.Printf("%s\n", y)
}
