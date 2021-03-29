package main

import (
   "encoding/json"
   "os"
)

func main() {
   s := "west & east"
   json.NewEncoder(os.Stdout).Encode(s)
}
