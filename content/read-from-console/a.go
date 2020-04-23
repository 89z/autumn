package main
import (
   "bufio"
   "os"
)
func main() {
   o1 := bufio.NewScanner(os.Stdin)
   o1.Scan()
   s1 := o1.Text()
   println(s1)
}
