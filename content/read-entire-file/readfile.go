package main
import "fmt"
import "io/ioutil"
func main() {
   y1, _ := ioutil.ReadFile("a.txt")
   fmt.Printf("%s", y1)
}
