package main
import "fmt"

func main() {
   var n = 1 + 2
   var s string

   switch n {
   case 5:
      s = "five"
   case 4, 3:
      s = "four or three"
   default:
      s = fmt.Sprint(n)
   }

   fmt.Println(s == "four or three")
}
