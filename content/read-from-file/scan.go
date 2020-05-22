package main
import (
   "bufio"
   "log"
   "os"
)
func main() {
   o1, e1 := os.Open("a.txt")
   if e1 != nil {
      log.Fatal(e1)
   }
   o2 := bufio.NewScanner(o1)
   a1 := []string{}
   for o2.Scan() {
      s1 := o2.Text()
      a1 = append(a1, s1)
   }
   log.Print(a1)
}
