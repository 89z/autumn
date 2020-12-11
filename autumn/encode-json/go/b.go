package main

import (
   "encoding/json"
   "os"
)

func main() {
   json.NewEncoder(os.Stdout).Encode(a)
}
