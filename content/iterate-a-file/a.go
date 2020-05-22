package main
import (
   "bufio"
   "log"
   "os"
)
func main() {
   o1, e := os.Open("a.txt")
   if e != nil {
      log.Fatal(e)
   }
   o2 := bufio.NewScanner(o1)
   for o2.Scan() {
      s1 := o2.Text()
      println(s1)
   }
}
