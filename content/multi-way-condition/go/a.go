package main
import "fmt"

func main() {
   var n = 46 / 10
   var s string

   switch n {
   case 7:
      s = "seven"
   case 6, 5:
      s = "six or five"
   default:
      s = fmt.Sprint(n)
   }

   fmt.Println(s == "4")
}
