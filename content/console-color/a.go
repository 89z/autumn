package main
import "github.com/multiverse-os/color" // 15 deps
func main() {
   s1 := "Sunday"
   // example 1
   s2 := color.Green(s1)
   // example 2
   s3 := color.Red(s1)
   // print
   println(s2, s3, s1)
}
