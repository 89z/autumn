package main

import (
   "fmt"
   "strings"
)

func main() {
   for _, t := range []struct{s, sep string} {
      {"", ""},             // []
      {"", ","},            // [""]
      {"north", ""},        // ["n" "o" "r" "t" "h"]
      {"south,north", ","}, // ["south" "north"]
   } {
      a := strings.Split(t.s, t.sep)
      fmt.Printf("%q\n", a)
   }
}
