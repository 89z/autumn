package main

import (
   "bufio"
   "bytes"
)

func scanner(s string, sep rune) *bufio.Scanner {
   scan := bufio.NewScanner(bytes.NewBufferString(s))
   scan.Split(func(s []byte, b bool) (int, []byte, error) {
      if b { return 0, nil, nil }
      n := bytes.IndexRune(s, sep)
      if n < 0 {
         return len(s), s, nil
      }
      return n + 1, s[:n], nil
   })
   return scan
}

func main() {
   s := scanner("north,east,south,west", ',')
   for s.Scan() {
      println(s.Text())
   }
}
