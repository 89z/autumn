package main

import (
   "encoding/json"
   "os"
)

func main() {
   a := []string{"/", "📗"}
   json.NewEncoder(os.Stdout).Encode(a)
}
