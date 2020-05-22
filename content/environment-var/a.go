package main
import "os"
func main() {
   // example 1
   s1 := os.Getenv("BROWSER")
   // example 2
   s2, b1 := os.LookupEnv("BROWSER")
   // print
   println(s1, s2, b1)
}
