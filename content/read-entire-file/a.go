package main
import (
   "bufio"
   "fmt"
   "os"
)
func main() {
   r1, _ := os.Open("a.txt")
   r2 := bufio.NewScanner(r1)
   var a1 []string
   for r2.Scan() {
      a1 = append(a1, r2.Text())
   }
   fmt.Printf("%q\n", a1)
}
