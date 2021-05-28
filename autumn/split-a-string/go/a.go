package main

import (
   "bufio"
   "strings"
)

func main() {
   r := strings.NewReader("North\nEast\nSouth\nWest")
   s := bufio.NewScanner(r)
   for s.Scan() {
      println(s.Text())
   }
}
