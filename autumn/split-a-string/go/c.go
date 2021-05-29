package main

import (
   "fmt"
   "strings"
)

type comma struct { tok string }

func (c *comma) Scan(state fmt.ScanState, verb rune) error {
   tok, err := state.Token(false, func(r rune) bool {
      return r != ','
   })
   if err != nil {
      return err
   }
   if _, _, err := state.ReadRune(); err != nil {
      if len(tok) == 0 {
         return err
      }
   }
   c.tok = string(tok)
   return nil
}

func main() {
   r := strings.NewReader("north,east,south,west")
   for {
      var c comma
      _, err := fmt.Fscan(r, &c)
      if err != nil { break }
      println(c.tok)
   }
}
