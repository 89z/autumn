package main

import (
   "bufio"
   "strings"
)

func comma(b []byte, eof bool) (int, []byte, error) {
   if eof { return 0, nil, nil }
   for n := range b {
      if b[n] == ',' {
         return n+1, b[:n], nil
      }
   }
   return len(b), b, nil
}

func main() {
   r := strings.NewReader("north,east,south,west")
   s := bufio.NewScanner(r)
   s.Split(comma)
   for s.Scan() {
      println(s.Text())
   }
}
