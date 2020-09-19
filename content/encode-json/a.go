package main

import (
   "encoding/json"
   "fmt"
   "os"
)

func main() {
   a := []string{"/", "📗"}
   y, e := json.Marshal(a)
   if e != nil {
      os.Exit(1)
   }
   fmt.Printf("%s\n", y)
}
