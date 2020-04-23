package main
import (
   "bufio"
   "os"
)
func main() {
   // example 1
   o1, _ := os.Open("a.txt")
   o2 := bufio.NewScanner(o1)
   o2.Scan()
   s1 := o2.Text()
   print(s1, "\n")
   // example 2
   o3, _ := os.Open("a.txt")
   o4 := bufio.NewReader(o3)
   s2, _ := o4.ReadString('\n')
   print(s2)
}
