package main
import (
   "bufio"
   "strings"
)
func main() {
   s1 := `Sunday
Monday`
   o1 := bufio.NewScanner(strings.NewReader(s1))
   for o1.Scan() {
      println(o1.Text())
   }
}
