package main

import (
   "fmt"
   "strings"
)

func main() {
   for _, t := range []struct{st, sep string} {
      {"", ""},             // []
      {"", ","},            // [""]
      {"north", ""},        // ["n" "o" "r" "t" "h"]
      {"south,north", ","}, // ["south" "north"]
   } {
      s := strings.Split(t.st, t.sep)
      fmt.Printf("%q\n", s)
   }
}
