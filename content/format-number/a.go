package main
import (
   "fmt"
   "golang.org/x/text/language"
   "golang.org/x/text/message"
)
func main() {
   n1 := 1000
   // example 1
   fmt.Printf("%5d\n", n1)
   // example 2
   p1 := message.NewPrinter(language.English)
   p1.Println(n1)
}
