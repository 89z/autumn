package main

import (
   "strings"
   "text/scanner"
)

func main() {
   var (
      r = strings.NewReader("north,east,south,west")
      s scanner.Scanner
   )
   s.Init(r)
   s.Whitespace = 1 << ','
   for s.Scan() != scanner.EOF {
      text := s.TokenText()
      println(text)
   }
}
