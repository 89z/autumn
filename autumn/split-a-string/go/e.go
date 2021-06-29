package main

import (
   "bufio"
   "strings"
)

func comma(data []byte, eof bool) (int, []byte, error) {
   var tok []byte
   for i, t := range data {
      if t != ',' {
         tok = append(tok, t)
      } else if tok != nil {
         return i+1, tok, nil
      }
   }
   return 0, nil, nil
}

func main() {
   s := bufio.NewScanner(strings.NewReader(",north,,south,"))
   s.Split(comma)
   for s.Scan() {
      println(s.Text())
   }
}
