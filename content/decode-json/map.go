package main

import (
   "encoding/json"
   "fmt"
   "log"
   "os"
)

func main() {
   // file to map
   o, e := os.Open("a.json")
   if e != nil {
      log.Fatal(e)
   }
   a := []map[string]string{}
   json.NewDecoder(o).Decode(&a)
   // string to map
   s := `[
      {"city":"Dresden","country":"Germany"},
      {"city":"Ostrava","country":"Czech Republic"}
   ]`
   y := []byte(s)
   a2 := []map[string]string{}
   json.Unmarshal(y, &a2)
   // print
   fmt.Print(a, "\n", a2, "\n")
}
