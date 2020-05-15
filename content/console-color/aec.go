package main
import (
   "fmt"
   "github.com/morikuni/aec"
)
func main() {
   s1 := "Sunday"
   // example 1
   fmt.Print(aec.GreenF)
   fmt.Print(s1)
   fmt.Print(aec.RedF)
   fmt.Print(s1)
   fmt.Print(aec.Reset)
   fmt.Println(s1)
   // example 2
   s2 := aec.GreenF.Apply(s1)
   s3 := aec.RedF.Apply(s1)
   s4 := aec.Reset + s1
   fmt.Println(s2 + s3 + s4)
}
