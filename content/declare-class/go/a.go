package main
import "fmt"

type Date struct {
   Year, Month int
}

func main() {
   // example 1
   o1 := Date{2019, 12}
   // example 2
   o2 := Date{Year: 2019}
   // print
   fmt.Println(o1, o2)
}
