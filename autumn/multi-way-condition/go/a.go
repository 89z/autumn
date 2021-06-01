package main
import "os"

func main() {
   n := len(os.Args)
   // example 1
   switch n {
   case 1:
      println("one")
   case 2, 3:
      println("some")
   default:
      println("more")
   }
   // example 2
   switch {
   case n == 1:
      println("one")
   case n == 2, n == 3:
      println("some")
   default:
      println("more")
   }
}
