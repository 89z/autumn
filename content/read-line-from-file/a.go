package main
import (
   "bufio"
   "os"
)
func main() {
   // example 1
   r1, _ := os.Open("a.txt")
   r2 := bufio.NewScanner(r1)
   r2.Scan()
   s1 := r2.Text()
   // example 2
   r3, _ := os.Open("a.txt")
   r4 := bufio.NewReader(r3)
   s2, _ := r4.ReadString('\n')
   // print
   print(s1, "\n")
   print(s2)
}
