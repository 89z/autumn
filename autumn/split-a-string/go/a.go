package main

import (
   "bufio"
   "strings"
)

const s = `north
east
south
west
`

func main() {
   b := bufio.NewScanner(strings.NewReader(s))
   for b.Scan() {
      println(b.Text())
   }
}
