package main
import "fmt"
import "io/ioutil"
func main() {
   a1, _ := ioutil.ReadFile("a.txt")
   fmt.Printf("%s", a1)
}
