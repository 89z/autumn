package main

import (
   "bufio"
   "strings"
)

func ScanNull(y []byte, b bool) (int, []byte, error) {
   n1 := 0
   for n1 < len(y) {
      if y[n1] != '\x00' {
         break
      }
      n1++
   }
   for n2 := n1; n2 < len(y); n2++ {
      if y[n2] == '\x00' {
         return n2 + 1, y[n1:n2], nil
      }
   }
   if b && len(y) > n1 {
      return len(y), y[n1:], nil
   }
   return n1, nil, nil
}

func main() {
   const input = "March\x00April\x00May\x00\x00June"
   scanner := bufio.NewScanner(strings.NewReader(input))
   scanner.Split(ScanNull)
   for scanner.Scan() {
      println(scanner.Text())
   }
}
