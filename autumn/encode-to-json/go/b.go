package main

import (
   "encoding/json"
   "io"
   "os"
)

func encode(s string, w io.Writer) error {
   enc := json.NewEncoder(w)
   enc.SetEscapeHTML(false)
   return enc.Encode(s)
}

func main() {
   encode("west & east", os.Stdout)
}
