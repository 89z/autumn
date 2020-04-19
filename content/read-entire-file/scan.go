package main
import "bufio"
import "fmt"
import "os"
func main() {
   // head
   o1, _ := os.Open("a.txt")
   o2 := bufio.NewScanner(o1)
   a1 := []string{}
   // body
   for o2.Scan() {
      s1 := o2.Text()
      a1 = append(a1, s1)
   }
   // foot
   fmt.Println(a1)
}
