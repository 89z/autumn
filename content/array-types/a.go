package main
import "fmt"
func main() {
   // example 1
   a1 := []string{"Sun", "Mon"}
   // example 2
   a2 := []interface{}{"Sun", "Mon"}
   // print
   fmt.Println(a1, a2)
}
