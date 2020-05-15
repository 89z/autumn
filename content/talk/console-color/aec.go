package main
import (
   "fmt"
   "github.com/morikuni/aec"
)
func main() {
   // example 1
   fmt.Print(aec.GreenF)
   fmt.Print("Sun")
   fmt.Print(aec.RedF)
   fmt.Print("Mon")
   fmt.Print(aec.Reset)
   fmt.Println("Tue")
   // example 2
   s_in := aec.GreenF.Apply("Sun") + aec.RedF.Apply("Mon") + aec.Reset + "Tue"
   fmt.Println(s_in)
}
