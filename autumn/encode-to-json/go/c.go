package main

import (
   "encoding/json"
   "os"
)

func main() {
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode([]int{10, 11})
}
