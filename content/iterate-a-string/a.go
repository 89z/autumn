package main
import (
   "bufio"
   "strings"
)
func main() {
   s1 := `sun mon tue
wed thu fri
`
   o1 := bufio.NewScanner(strings.NewReader(s1))
   for o1.Scan() {
      println(o1.Text())
   }
}
