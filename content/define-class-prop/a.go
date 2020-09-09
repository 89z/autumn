package main
import "fmt"
func main() {

   // example 1
   type May struct {
      Day int
      Weekday string
   }
   o := May{8, "Tuesday"}

   // example 2
   type June struct {
      int
      string
   }
   o2 := June{8, "Tuesday"}

   // example 3
   type July struct {
      Day int
      Weekday string
   }
   o3 := July{Day: 8}

   // print
   fmt.Println(o, o2, o3)
}
