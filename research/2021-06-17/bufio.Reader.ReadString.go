package main

import (
   "bufio"
   "strings"
)

func comma(r *bufio.Reader) (string, error) {
   s, e := r.ReadString(',')
   if s == "" {
      return "", e
   }
   if e != nil {
      return s, nil
   }
   return s[:len(s)-1], nil
}

func main() {
   r := bufio.NewReader(strings.NewReader("north,east,south,west"))
   for {
      s, e := comma(r)
      if e != nil { break }
      println(s)
   }
}
