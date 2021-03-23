package main

import (
   "encoding/json"
   "log"
   "os"
)

func main() {
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   err := enc.Encode("May & June")
   if err != nil {
      log.Fatal(err)
   }
}
