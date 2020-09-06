package main

import (
   "encoding/json"
   "fmt"
   "log"
   "os"
)

type Place struct {
   City string
   Country string
}

func main() {
   // file to object
   o, e := os.Open("a.json")
   if e != nil {
      log.Fatal(e)
   }
   a := []Place{}
   json.NewDecoder(o).Decode(&a)
   // string to object
   s := `[
      {"city":"Dresden","country":"Germany"},
      {"city":"Ostrava","country":"Czech Republic"}
   ]`
   y := []byte(s)
   a2 := []Place{}
   json.Unmarshal(y, &a2)
   // print
   fmt.Printf("%+v\n%+v\n", a, a2)
}
