package main
import "fmt"
func main() {
   // example 1
   var a1 = []string{"Sun"}
   a1 = append(a1, "Mon")
   // example 2
   var a2 = []string{"Sun"}
   a2 = append(a2, []string{"Mon"}...)
   // print
   fmt.Println(a1, a2)
}
