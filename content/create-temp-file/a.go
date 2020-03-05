package main
import "io/ioutil"
func main() {
   // create only
   f1, _ := ioutil.TempFile("", "")
   s1 := f1.Name()
   // print
   println(s1)
}
