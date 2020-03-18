package main
import (
   "os"
   "bufio"
)
func main() {
   r1, _ := os.Create("a.txt")
   r1.WriteString("Sunday\n")
   r1.Seek(0, 0)
   r2 := bufio.NewScanner(r1)
   r2.Scan()
   s1 := r2.Text()
   println(s1)
}
