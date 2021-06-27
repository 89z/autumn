package main

import (
   "fmt"
   "io"
   "strings"
)

func scan(s io.Reader, sep rune) []rune {
   var (
      a rune
      b []rune
   )
   for {
      _, err := fmt.Fscanf(s, "%c", &a)
      if err != nil {
         break
      } else if a != sep {
         b = append(b, a)
      } else if b != nil {
         break
      }
   }
   return b
}

func main() {
   s := strings.NewReader(",north,,south,")
   for {
      text := scan(s, ',')
      if text == nil {
         break
      }
      fmt.Printf("%c\n", text)
   }
}
