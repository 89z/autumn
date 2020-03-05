package main
import "os"
func main() {
   // example 1
   r1, _ := os.Open("a.txt")
   // example 2
   r2, _ := os.Create("a.txt")
   // print
   println(r1, r2)
}
