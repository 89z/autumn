package main
import "os"
func main() {
   s1 := os.Getenv("HOME")
   println(s1)
}
