package main

import (
   "encoding/json"
   "os"
)

func main() {
   o := json.NewEncoder(os.Stdout)
   o.SetIndent("", "\t")
   o.Encode(a)
}
