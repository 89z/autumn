package main
import "fmt"

func main() {
   type Date struct {
      Year, Month, Day int
   }
   // example 1
   o1 := Date{2019, 12, 31}
   // example 2
   o2 := Date{Year: 2019}
   // print
   fmt.Println(o1, o2)
}
