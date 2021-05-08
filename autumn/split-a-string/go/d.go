package main

import (
   "fmt"
   "strings"
)

func main() {
   r := strings.NewReader("north east\nsouth west\n")
   for {
      var s, t string
      _, e := fmt.Fscanln(r, &s, &t)
      if e != nil { break }
      fmt.Printf("%q %q\n", s, t)
   }
}
