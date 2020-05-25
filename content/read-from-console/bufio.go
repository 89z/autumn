package main
import (
   "bufio"
   "os"
)
func main() {
   o := bufio.NewScanner(os.Stdin)
   o.Scan()
   s := o.Text()
   println(s)
}
