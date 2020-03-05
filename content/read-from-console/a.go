package main
import (
   "bufio"
   "os"
)
func main() {
   r1 := bufio.NewScanner(os.Stdin)
   r1.Scan()
   s1 := r1.Text()
   println(s1)
}
