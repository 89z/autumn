package main

import (
   "bufio"
   "bytes"
)

func comma(s []byte, b bool) (int, []byte, error) {
   n := bytes.IndexRune(s, ',')
   switch {
   case b:     return 0,      nil,   nil
   case n < 0: return len(s), s,     nil
   default:    return n + 1,  s[:n], nil
   }
}

func main() {
   s := bufio.NewScanner(bytes.NewBufferString("north,east,south,west"))
   s.Split(comma)
   for s.Scan() {
      println(s.Text())
   }
}
