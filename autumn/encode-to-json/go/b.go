package main

import (
   "encoding/json"
   "io"
   "os"
)

func encode(s string, w io.Writer) error {
   e := json.NewEncoder(w)
   e.SetEscapeHTML(false)
   return e.Encode(s)
}

func main() {
   encode("west & east", os.Stdout)
}
