package main

import (
   "strings"
   "text/scanner"
)

type test struct {
   s string
   sep rune
}

func main() {
   tests := []test{
      {"Wonder.Woman.1984.2020.IMAX", '.'}, {"north,east,south,west", ','},
   }
   for _, t := range tests {
      var (
         r = strings.NewReader(t.s)
         s scanner.Scanner
      )
      s.Init(r)
      s.Whitespace = 1 << t.sep // max is '?'
      for s.Scan() != scanner.EOF {
         text := s.TokenText()
         println(text)
      }
   }
}
