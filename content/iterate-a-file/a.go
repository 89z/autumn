package main
import (
   "bufio"
   "log"
   "os"
)
func main() bool {
   o1, e1 := os.Open("a.txt")
   if e1 != nil {
      log.Fatal(e1)
   }
   o2 := bufio.NewScanner(o1)
   for o2.Scan() {
      s1 := o2.Text()
      println(s1)
   }
}
