package main
import (
   "bufio"
   "os"
)
func main() {
   o1, _ := os.Open("a.txt")
   o2 := bufio.NewScanner(o1)
   o2.Scan()
   s1 := o2.Text()
   println(s1)
}
