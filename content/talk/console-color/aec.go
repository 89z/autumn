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
   s_in := aec.GreenF.Apply(s1) + aec.RedF.Apply(s1) + aec.Reset + s1
   fmt.Println(s_in)
}
