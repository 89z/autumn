package main
import (
   "bufio"
   "os"
)
func main() {
   o1, e := os.Open("a.txt")
   if e != nil {
      os.Exit(1)
   }
   o2 := bufio.NewScanner(o1)
   for o2.Scan() {
      s := o2.Text()
      println(s)
   }
}
