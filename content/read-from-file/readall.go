package main
import "fmt"
import "io/ioutil"
import "os"
func main() {
   o1, _ := os.Open("a.txt")
   y1, _ := ioutil.ReadAll(o1)
   fmt.Printf("%s", y1)
}
