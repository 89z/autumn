package main
import "fmt"

func main() {
   type Date struct {
      Year, Month, Day int
   }
   // example 1
   o := Date{2020, 9, 10}
   // example 2
   o2 := Date{Year: 2020}
   // print
   fmt.Println(o, o2)
}
