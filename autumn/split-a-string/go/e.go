package main

import (
   "bufio"
   "strings"
)

func comma(data []byte, eof bool) (int, []byte, error) {
   a := -1
   for b, c := range data {
      if c == ',' {
         if a >= 0 {
            return b+1, data[a:b], nil
         }
      } else if a < 0 {
         a = b
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
