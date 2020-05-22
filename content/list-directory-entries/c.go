package main
import (
   "fmt"
   "log"
   "path/filepath"
)
func main() {
   a, e := filepath.Glob("*")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
