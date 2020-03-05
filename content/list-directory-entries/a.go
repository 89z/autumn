package main
import "io/ioutil"
func main() {
   a1, _ := ioutil.ReadDir(".")
   for _, o1 := range a1 {
      s1 := o1.Name()
      println(s1)
   }
}
